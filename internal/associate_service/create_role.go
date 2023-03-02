package associate_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) CreateRole(ctx context.Context, req *desc.CreateRoleRequest) (*desc.CreateRoleResponse, error) {
	access, err := s.checkUserAdmin(ctx)
	if err != nil {
		return nil, err
	}
	if !access {
		return nil, status.Error(codes.PermissionDenied, "only for admins")
	}
	res, err := s.roleService.CreateRole(ctx, s.protoCreateRoleRequestToRole(req))
	if err != nil {
		return nil, err
	}
	return s.roleToProtoCreateRoleResponse(res), nil
}

func (s *AssociateApiServiceServer) protoCreateRoleRequestToRole(req *desc.CreateRoleRequest) datastruct.Role {
	return datastruct.Role{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}
}

func (s *AssociateApiServiceServer) roleToProtoCreateRoleResponse(resp *datastruct.Role) *desc.CreateRoleResponse {
	return &desc.CreateRoleResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description,
	}
}
