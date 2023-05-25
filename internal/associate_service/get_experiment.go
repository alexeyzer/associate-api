package associate_service

import (
	"context"

	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) GetExperiment(ctx context.Context, req *desc.GetExperimentRequest) (*desc.GetExperimentResponse, error) {
	_, err := s.CheckAutorize(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.experimentService.GetExperiment(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetExperimentResponse{
		Id: res.ID,
	}, nil
}
