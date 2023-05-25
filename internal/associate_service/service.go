package associate_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/alexeyzer/associate-api/config"
	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/alexeyzer/associate-api/internal/pkg/service"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

type AssociateApiServiceServer struct {
	userService       service.UserService
	roleService       service.RoleService
	userRoleService   service.UserRoleService
	experimentService service.ExperimentService
	stimusWordService service.StimusWordService
	desc.UnimplementedAssociateApiServiceServer
}

const adminRole = "admin"

func (s *AssociateApiServiceServer) checkUserAdmin(ctx context.Context) (bool, error) {
	sessionId := s.GetSessionIDFromContext(ctx)
	if sessionId == "" {
		return false, status.Errorf(codes.Unauthenticated, "sessionID does`t exists, please, login")
	}
	res, err := s.userService.SessionCheck(ctx, sessionId)
	if err != nil {
		return false, err
	}
	for _, role := range res.Roles {
		if role == adminRole {
			return true, nil
		}
	}
	return false, nil
}

func (s *AssociateApiServiceServer) CheckAutorize(ctx context.Context) (*datastruct.UserWithRoles, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		val := md.Get(config.SessionKey)
		if len(val) > 0 {
			res, err := s.userService.SessionCheck(ctx, val[0])
			if err != nil {
				return nil, err
			}
			return res, nil
		}
		log.Info("no value with key:", config.SessionKey)
	}
	log.Info("no metadata")
	return nil, status.Error(codes.PermissionDenied, "login pls")
}

func (s *AssociateApiServiceServer) GetSessionIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		val := md.Get(config.SessionKey)
		if len(val) > 0 {
			return val[0]
		}
		log.Info("no value with key:", config.SessionKey)
	}
	log.Info("no metadata")
	return ""
}

func NewAssociateApiServiceServer(
	userService service.UserService,
	roleService service.RoleService,
	userRoleService service.UserRoleService,
	experimentService service.ExperimentService,
	stimusWordService service.StimusWordService,
) *AssociateApiServiceServer {
	return &AssociateApiServiceServer{
		userService:       userService,
		roleService:       roleService,
		userRoleService:   userRoleService,
		experimentService: experimentService,
		stimusWordService: stimusWordService,
	}
}
