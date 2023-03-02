package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AssociateApiServiceServer) SessionCheck(ctx context.Context, _ *emptypb.Empty) (*desc.SessionCheckResponse, error) {
	sessionId := s.GetSessionIDFromContext(ctx)
	if sessionId == "" {
		return nil, status.Errorf(codes.Unauthenticated, "sessionID does`t exists, please, login")
	}
	resp, err := s.userService.SessionCheck(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	return s.userWithRolesToProtoSessionCheckResponse(resp), nil
}

func (s *AssociateApiServiceServer) userWithRolesToProtoSessionCheckResponse(resp *datastruct.UserWithRoles) *desc.SessionCheckResponse {
	internalResp := &desc.SessionCheckResponse{
		UserId:    resp.ID,
		Email:     resp.Email,
		UserRoles: resp.Roles,
	}
	return internalResp
}
