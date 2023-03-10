package associate_service

import (
	"context"

	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	roles, sessionID, res, err := s.userService.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := &desc.LoginResponse{
		Id:                 res.ID,
		Name:               res.Name,
		Surname:            res.Surname,
		Patronymic:         res.Patronymic,
		Phone:              res.Phone,
		Email:              res.Email,
		Session:            sessionID,
		AccessToAdminPanel: len(roles) > 0,
		Roles:              make([]*desc.CreateRoleResponse, 0, len(roles)),
	}
	for _, role := range roles {
		resp.Roles = append(resp.Roles, &desc.CreateRoleResponse{
			Id:          role.RoleID,
			Name:        role.RoleName,
			Description: "",
		})
	}
	return resp, nil
}
