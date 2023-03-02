package associate_service

import (
	"context"

	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AssociateApiServiceServer) DeleteRole(ctx context.Context, req *desc.DeleteRoleRequest) (*emptypb.Empty, error) {
	access, err := s.checkUserAdmin(ctx)
	if err != nil {
		return nil, err
	}
	if !access {
		return nil, status.Error(codes.PermissionDenied, "only for admins")
	}

	err = s.roleService.DeleteRole(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
