package service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/alexeyzer/associate-api/internal/pkg/repository"
)

type ExperimentResultService interface {
	CreateExperimentResult(ctx context.Context, req []*datastruct.ExperimentResult) ([]*datastruct.ExperimentResultResp, error)
	CalucateExperiment(ctx context.Context, experimentID int64) error
	GetCalculatedResulsts(ctx context.Context, experimentID int64, number, limit int64, names []string) ([]*datastruct.ExperimentResultCalculated, error)
}

type experimentResultService struct {
	dao repository.DAO
}

func (e *experimentResultService) GetCalculatedResulsts(ctx context.Context, experimentID int64, number, limit int64, names []string) ([]*datastruct.ExperimentResultCalculated, error) {
	return e.dao.ExperimentResultCalculatedQuery().GetCalculatedResulsts(ctx, 0, []int64{experimentID}, number, limit, names)
}

func (e *experimentResultService) CalucateExperiment(ctx context.Context, experimentID int64) error {
	results, err := e.dao.ExperimentResultQuery().List(ctx, 0, []int64{experimentID}, 0, 0, []string{})
	if err != nil {
		return err
	}
	grahp := make(map[int64]map[int64]int64, len(results))

	for _, r := range results {
		_, ok := grahp[r.StimusWordID]
		if !ok {
			grahp[r.StimusWordID] = make(map[int64]int64, 30)
		}
		grahp[r.StimusWordID][r.AssotiationWordID] = grahp[r.StimusWordID][r.AssotiationWordID] + 1
	}
	result := make([]*datastruct.ExperimentResultCalculated, 0, len(grahp))

	for key, mapp := range grahp {
		for associate, amount := range mapp {
			result = append(result, &datastruct.ExperimentResultCalculated{
				ExperimentID:      experimentID,
				StimusWordID:      key,
				AssotiationWordID: associate,
				Amount:            amount,
			})
		}
	}
	return e.dao.ExperimentResultCalculatedQuery().BatchCreate(ctx, result)
}

func (e *experimentResultService) CreateExperimentResult(ctx context.Context, req []*datastruct.ExperimentResult) ([]*datastruct.ExperimentResultResp, error) {
	for _, r := range req {
		res, err := e.dao.AssociateWordQuery().GetByName(ctx, r.AssotiationWord)
		if err == nil {
			r.AssotiationWordID = res.ID
			continue
		}

		res, err = e.dao.AssociateWordQuery().Create(ctx, datastruct.AssociateWord{
			Name: r.AssotiationWord,
		})
		if err != nil {
			return nil, err
		}
		r.AssotiationWordID = res.ID
	}

	repositoryes, err := e.dao.ExperimentResultQuery().BatchCreate(ctx, req)
	if err != nil {
		return nil, err
	}

	err = e.CalucateExperiment(ctx, req[0].ExperimentID)
	if err != nil {
		return nil, err
	}
	return repositoryes, nil
}

func NewExperimentResultService(dao repository.DAO) ExperimentResultService {
	return &experimentResultService{dao: dao}
}
