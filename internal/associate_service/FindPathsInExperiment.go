package associate_service

import (
	"context"

	pb "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) FindPathsInExperiment(ctx context.Context, req *pb.FindPathsInExperimentRequest) (*pb.FindPathsInExperimentResponse, error) {
	_, err := s.CheckAutorize(ctx)
	if err != nil {
		return nil, err
	}

	result, err := s.experimentResultService.FindPathsInExperiment(ctx, req.GetId(), req.GetWord1(), req.GetWord2())
	if err != nil {
		return nil, err
	}

	return &pb.FindPathsInExperimentResponse{
		Paths: BuildChain(result),
	}, nil
}
