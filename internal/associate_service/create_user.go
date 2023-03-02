package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AssociateApiServiceServer) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	sessionId := s.GetSessionIDFromContext(ctx)
	if sessionId != "" {
		return nil, status.Errorf(codes.InvalidArgument, "already logged in")
	}
	res, err := s.userService.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.serviceCreateUserResponseToProtoCreateUserResponse(res), nil
}

func (s *AssociateApiServiceServer) serviceCreateUserResponseToProtoCreateUserResponse(res *datastruct.User) *desc.CreateUserResponse {
	return &desc.CreateUserResponse{
		Id:         res.ID,
		Name:       res.Name,
		Surname:    res.Surname,
		Patronymic: res.Patronymic,
		Phone:      res.Phone,
		Email:      res.Email,
	}
}
