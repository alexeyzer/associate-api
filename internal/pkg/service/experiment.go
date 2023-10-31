package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/alexeyzer/associate-api/internal/pkg/repository"
)

type ExperimentService interface {
	CreateExperiment(ctx context.Context, req datastruct.Experiment) (*datastruct.Experiment, error)
	GetExperiment(ctx context.Context, id int64) (*datastruct.Experiment, error)
	ListExperiment(ctx context.Context, number, limit int64) ([]*datastruct.Experiment, error)
}

type experimentService struct {
	dao repository.DAO
}

func (e *experimentService) ListExperiment(ctx context.Context, number, limit int64) ([]*datastruct.Experiment, error) {
	return e.dao.ExperimentQuery().List(ctx, number, limit)
}

func (e *experimentService) CreateExperiment(ctx context.Context, req datastruct.Experiment) (*datastruct.Experiment, error) {
	for _, stimus := range req.ExperimentStimuses {
		if stimus.StimusID != 0 {
			_, err := e.dao.StimusWordQuery().GetByID(ctx, stimus.StimusID)
			if err != nil {
				return nil, err
			}
		}
		if stimus.Name != "" {

			res, err := e.dao.StimusWordQuery().Create(ctx, datastruct.StimusWord{
				Name: stimus.Name,
			})
			if err != nil {
				return nil, err
			}
			stimus.StimusID = res.ID
		}
	}

	exists, err := e.dao.ExperimentQuery().Exists(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, status.Errorf(codes.InvalidArgument, "experiment with name %s already exists", req.Name)
	}

	res, err := e.dao.ExperimentQuery().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *experimentService) GetExperiment(ctx context.Context, id int64) (*datastruct.Experiment, error) {
	experiment, err := e.dao.ExperimentQuery().GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return experiment, nil
}

func NewExperimentService(dao repository.DAO) ExperimentService {
	return &experimentService{dao: dao}
}
