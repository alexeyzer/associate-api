package associate_service

import (
	"context"
	"fmt"

	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) ListExperiment(ctx context.Context, req *desc.ListExperimentRequest) (*desc.ListExperimentResponse, error) {
	var userID int64

	user, err := s.CheckAutorize(ctx)
	if err == nil {
		userID = user.ID
	}
	if !req.UserExperiments {
		userID = 0
	}

	fmt.Println(userID)
	var name *string

	if req.GetName() != nil && req.GetName().GetValue() != "" {
		nameValue := req.GetName().GetValue()
		name = &nameValue
	}

	res, err := s.experimentService.ListExperiment(ctx, req.GetPage().GetNumber(), req.GetPage().GetLimit(), userID, req.GetUserExperiments(), name)
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
