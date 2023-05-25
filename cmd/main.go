package main

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/alexeyzer/associate-api/config"
	"github.com/alexeyzer/associate-api/internal/associate_service"
	"github.com/alexeyzer/associate-api/internal/client"
	"github.com/alexeyzer/associate-api/internal/pkg/repository"
	"github.com/alexeyzer/associate-api/internal/pkg/service"
	gw "github.com/alexeyzer/associate-api/pb/api/associate/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func serveSwagger(mux *http.ServeMux) {
	prefix := "/swagger/"
	sh := http.StripPrefix(prefix, http.FileServer(http.Dir("./swagger/")))
	mux.Handle(prefix, sh)
}

// look up session and pass sessionId in to context if it exists
func gatewayMetadataAnnotator(_ context.Context, r *http.Request) metadata.MD {
	SessionID := r.Header.Get(config.SessionKey)
	c, err := r.Cookie(config.SessionKey)
	if err == nil {
		SessionID = c.Value
	}
	if SessionID != "" {
		return metadata.Pairs(config.SessionKey, SessionID)
	}
	return metadata.Pairs()
}

func httpResponseModifier(ctx context.Context, w http.ResponseWriter, _ proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	sessionID := md.HeaderMD.Get(config.SessionKey)
	logout := md.HeaderMD.Get(config.Config.Auth.LogoutKey)
	if len(sessionID) > 0 {
		if len(logout) == 0 {
			http.SetCookie(w, &http.Cookie{
				Name:     config.SessionKey,
				Value:    sessionID[0],
				Path:     "/",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour * 24),
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
			})
		} else {
			w.Header().Add(config.SessionKey, sessionID[0])
			http.SetCookie(w, &http.Cookie{
				Name:     config.SessionKey,
				Value:    sessionID[0],
				Path:     "/",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Duration(-1) * time.Hour * 24),
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
			})
		}
	}
	return nil
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "sessionid, Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		if (*r).Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func RunServer(ctx context.Context, associateApiServiceServer *associate_service.AssociateApiServiceServer) error {

	grpcLis, err := net.Listen("tcp", ":"+config.Config.App.GrpcPort)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_logrus.UnaryServerInterceptor(log.WithContext(ctx).WithTime(time.Time{})),
		grpc_validator.UnaryServerInterceptor(),
		grpc_prometheus.UnaryServerInterceptor)),
	)
	gw.RegisterAssociateApiServiceServer(grpcServer, associateApiServiceServer)

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux(runtime.WithMetadata(gatewayMetadataAnnotator), runtime.WithForwardResponseOption(httpResponseModifier))
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", cors(gwmux))
	serveSwagger(mux)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(),
	}
	err = gw.RegisterAssociateApiServiceHandlerFromEndpoint(ctx, gwmux, ":"+config.Config.App.GrpcPort, opts)
	if err != nil {
		return err
	}

	grpc_prometheus.Register(grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()

	go func() {
		err = grpcServer.Serve(grpcLis)
		log.Fatal(err)
	}()
	log.Println("app started")
	err = http.ListenAndServe(":"+config.Config.App.HttpPort, mux)
	log.Fatal(err)
	return err
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := config.ReadConf("./config/config.yaml")
	if err != nil {
		log.Fatal("Failed to create config: ", err)
	}

	dao, err := repository.NewDao()
	if err != nil {
		log.Fatalf("Failed to connect to db: %s", err)
	}

	redis, err := client.NewRedisClient(ctx)
	if err != nil {
		log.Fatal("Failed to connect to redis db: ", err)
	}

	userService := service.NewUserService(dao, redis)
	roleService := service.NewRoleService(dao)
	userRoleService := service.NewUserRoleService(dao)
	experimentService := service.NewExperimentService(dao)
	stimusWordService := service.NewStimusWordService(dao)

	associateApiServiceServer := associate_service.NewAssociateApiServiceServer(userService, roleService, userRoleService, experimentService, stimusWordService)
	if err := RunServer(ctx, associateApiServiceServer); err != nil {
		log.Fatal(err)
	}
}
