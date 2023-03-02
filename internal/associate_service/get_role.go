package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) GetRole(ctx context.Context, req *desc.GetRoleRequest) (*desc.GetRoleResponse, error) {
	res, err := s.roleService.GetRole(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.roleToProtoGetRoleResponse(res), nil
}

func (s *AssociateApiServiceServer) roleToProtoGetRoleResponse(resp *datastruct.Role) *desc.GetRoleResponse {
	return &desc.GetRoleResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description,
	}
}
