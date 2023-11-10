package service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/alexeyzer/associate-api/internal/pkg/repository"
)

type ExperimentResultService interface {
	CreateExperimentResult(ctx context.Context, req []*datastruct.ExperimentResult) ([]*datastruct.ExperimentResultResp, error)
}

type experimentResultService struct {
	dao repository.DAO
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
	return repositoryes, nil
}

func NewExperimentResultService(dao repository.DAO) ExperimentResultService {
	return &experimentResultService{dao: dao}
}
