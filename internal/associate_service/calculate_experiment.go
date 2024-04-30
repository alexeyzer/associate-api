package associate_service

import (
	"context"

	pb "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) CalculateExperiment(ctx context.Context, req *pb.CalculateExperimentRequest) (*pb.CalculateExperimentResponse, error) {
	err := s.experimentResultService.CalucateExperiment(ctx, req.GetExprimentId())
	if err != nil {
		return nil, err
	}
	return &pb.CalculateExperimentResponse{}, nil
}
