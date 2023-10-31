package associate_service

import (
	"context"

	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) ListExperiment(ctx context.Context, req *desc.ListExperimentRequest) (*desc.ListExperimentResponse, error) {

	res, err := s.experimentService.ListExperiment(ctx, req.GetPage().GetNumber(), req.GetPage().GetLimit())
	if err != nil {
		return nil, err
	}

	resp := &desc.ListExperimentResponse{
		Experiments: make([]*desc.GetExperimentResponse, 0, 1),
	}
	for _, experiment := range res {
		resp.Experiments = append(resp.Experiments, datastructExperimentToPB(experiment))
	}

	return resp, nil
}
