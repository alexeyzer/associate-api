package associate_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/alexeyzer/associate-api/config"
	"github.com/alexeyzer/associate-api/internal/pkg/service"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

type AssociateApiServiceServer struct {
	userService     service.UserService
	roleService     service.RoleService
	userRoleService service.UserRoleService
	desc.UnimplementedAssociateApiServiceServer
}

const adminRole = "admin"

func (s *AssociateApiServiceServer) checkUserAdmin(ctx context.Context) (bool, error) {
	res, err := s.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return false, err
	}
	for _, role := range res.UserRoles {
		if role == adminRole {
			return true, nil
		}
	}
	return false, nil
}

func (s *AssociateApiServiceServer) GetSessionIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		val := md.Get(config.Config.Auth.SessionKey)
		if len(val) > 0 {
			return val[0]
		}
		log.Info("no value with key:", config.Config.Auth.SessionKey)
	}
	log.Info("no metadata")
	return ""
}

func NewAssociateApiServiceServer(
	userService service.UserService,
	roleService service.RoleService,
	userRoleService service.UserRoleService,
) *AssociateApiServiceServer {
	return &AssociateApiServiceServer{
		userService:     userService,
		roleService:     roleService,
		userRoleService: userRoleService,
	}
}
