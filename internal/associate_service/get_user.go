package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"

	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	resp, roles, err := s.userService.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.toProtoGetUserResponse(resp, roles), nil
}

func (s *AssociateApiServiceServer) toProtoGetUserResponse(res *datastruct.User, roles []*datastruct.UserRoleWithName) *desc.GetUserResponse {
	resp := &desc.GetUserResponse{
		Id:         res.ID,
		Name:       res.Name,
		Surname:    res.Surname,
		Patronymic: res.Patronymic,
		Phone:      res.Phone,
		Email:      res.Email,
		Roles:      make([]*desc.GetUserResponse_UserRoles, 0, len(roles)),
	}
	for _, role := range roles {
		resp.Roles = append(resp.Roles, &desc.GetUserResponse_UserRoles{
			UserRoleId: role.ID,
			Name:       role.RoleName,
		})
	}
	return resp
}
