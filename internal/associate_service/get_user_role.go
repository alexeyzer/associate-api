package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) GetUserRole(ctx context.Context, req *desc.GetUserRoleRequest) (*desc.GetUserRoleResponse, error) {
	resp, err := s.userRoleService.GetUserRole(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.userRoleToProtoGetUserRoleResponse(resp), nil
}

func (s *AssociateApiServiceServer) userRoleToProtoGetUserRoleResponse(resp *datastruct.UserRole) *desc.GetUserRoleResponse {
	return &desc.GetUserRoleResponse{
		Id:     resp.ID,
		UserId: resp.UserID,
		RoleId: resp.RoleID,
	}
}
