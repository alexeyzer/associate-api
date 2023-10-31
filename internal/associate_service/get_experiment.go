package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

	return datastructExperimentToPB(res), nil
}

func datastructExperimentToPB(d *datastruct.Experiment) *desc.GetExperimentResponse {
	response := &desc.GetExperimentResponse{
		Id:                 d.ID,
		Name:               d.Name,
		ExperimentStimuses: make([]*desc.ExperimentStimuses, 0, len(d.ExperimentStimuses)),
	}
	for _, stimus := range d.ExperimentStimuses {
		response.ExperimentStimuses = append(response.ExperimentStimuses, ExperimentStimusTopb(stimus))
	}

	return response
}

func ExperimentStimusTopb(d *datastruct.ExperimentStimus) *desc.ExperimentStimuses {
	stimus := &desc.ExperimentStimuses{
		Id:   wrapperspb.Int64(d.StimusID),
		Name: wrapperspb.String(d.Name),
	}
	if d.LimitedResponseTime != nil {
		stimus.LimitedResponseTime = wrapperspb.Int64(*d.LimitedResponseTime)
	}
	return stimus
}
