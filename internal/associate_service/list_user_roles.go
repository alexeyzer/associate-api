package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"

	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) ListUserRoles(ctx context.Context, req *desc.ListUserRolesRequest) (*desc.ListUserRolesResponse, error) {
	resp, err := s.userRoleService.ListUserRoles(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return s.userRolesToProtoListUserRolesResponse(resp), nil
}

func (s *AssociateApiServiceServer) userRolesToProtoListUserRolesResponse(resp []*datastruct.UserRoleWithName) *desc.ListUserRolesResponse {
	internalResp := &desc.ListUserRolesResponse{
		UserRoles: make([]string, 0, len(resp)),
	}
	for _, item := range resp {
		internalResp.UserRoles = append(internalResp.UserRoles, item.RoleName)
	}
	return internalResp
}
