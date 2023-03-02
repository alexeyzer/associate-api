package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"

	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) UpdateRole(ctx context.Context, req *desc.UpdateRoleRequest) (*desc.UpdateRoleResponse, error) {
	resp, err := s.roleService.UpdateRole(ctx, datastruct.Role{
		ID:          req.GetId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, err
	}
	return &desc.UpdateRoleResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description,
	}, nil
}
