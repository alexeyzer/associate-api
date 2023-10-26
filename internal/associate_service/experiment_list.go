package associate_service

import (
	"context"

	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) ListExperiment(ctx context.Context, req *desc.ListExperimentRequest) (*desc.ListExperimentResponse, error) {
	_, err := s.CheckAutorize(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.experimentService.GetExperiment(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	resp := &desc.ListExperimentResponse{
		Experiments: make([]*desc.ListExperimentResponse_Experiment, 0, 1),
	}
	for _, experiment := range res {
	}

	return resp, nil
}
