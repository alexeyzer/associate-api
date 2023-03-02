package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AssociateApiServiceServer) ListUsers(ctx context.Context, _ *emptypb.Empty) (*desc.ListUsersResponse, error) {
	resp, err := s.userService.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return s.usersToProtoListUsersResponse(resp), nil
}

func (s *AssociateApiServiceServer) usersToProtoListUsersResponse(users []*datastruct.User) *desc.ListUsersResponse {
	resp := &desc.ListUsersResponse{
		Users: make([]*desc.CreateUserResponse, 0, len(users)),
	}

	for _, user := range users {
		resp.Users = append(resp.Users, s.serviceCreateUserResponseToProtoCreateUserResponse(user))
	}
	return resp
}
