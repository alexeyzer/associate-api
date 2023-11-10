package associate_service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/associate-api/pb/api/associate/v1"
)

func (s *AssociateApiServiceServer) CreateExperimentResult(ctx context.Context, req *desc.CreateExperimentResultRequest) (*desc.CreateExperimentResultResponse, error) {
	_, err := s.CheckAutorize(ctx)
	if err != nil {
		return nil, err
	}

	reqq := make([]*datastruct.ExperimentResult, 0, len(req.GetAnswers()))

	for _, r := range req.GetAnswers() {
		reqq = append(reqq, &datastruct.ExperimentResult{
			ExperimentID:    req.GetExperimentId(),
			UserID:          req.GetUserId(),
			SessionID:       req.GetSessionId(),
			StimusWordID:    r.GetStimusWordId(),
			AssotiationWord: r.GetAssotiationWord(),
			TimeSpend:       r.GetTimeSpend(),
		})
	}

	_, err = s.experimentResultService.CreateExperimentResult(ctx, reqq)
	if err != nil {
		return nil, err
	}
	return &desc.CreateExperimentResultResponse{}, nil
}
