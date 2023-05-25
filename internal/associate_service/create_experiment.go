package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) CreateExperiment(ctx context.Context, req *desc.CreateExperimentRequest) (*desc.CreateExperimentResponse, error) {
	user, err := s.CheckAutorize(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.experimentService.CreateExperiment(ctx, *toCreateExperiment(user.ID, req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateExperimentResponse{
		Id: res.ID,
	}, nil
}

func toCreateExperiment(id int64, req *desc.CreateExperimentRequest) *datastruct.Experiment {

	exp := &datastruct.Experiment{
		Name:               req.GetName(),
		CreatorID:          id,
		Status:             datastruct.ExperimentStatus_CREATED,
		ExperimentStimuses: make(datastruct.ExperimentStimuses, 0, len(req.GetExperimentStimuses())),
	}
	if req.GetDescription() != nil {
		exp.Description = &req.GetDescription().Value
	}
	if req.GetRequeiredAmount() != nil {
		exp.RequiredAmount = &req.GetRequeiredAmount().Value
	}
	for _, stimus := range req.GetExperimentStimuses() {
		stim := &datastruct.ExperimentStimus{}
		if stimus.GetId() != nil {
			stim.StimusID = stimus.GetId().GetValue()
		}
		if stimus.GetName() != nil {
			stim.Name = stimus.GetName().Value
		}
		if stimus.GetLimitedResponseTime() != nil {
			stim.LimitedResponseTime = &stimus.GetLimitedResponseTime().Value
		}
		exp.ExperimentStimuses = append(exp.ExperimentStimuses, stim)
	}
	return exp
}
