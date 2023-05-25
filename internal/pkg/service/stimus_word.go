package service

import (
	"context"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/alexeyzer/associate-api/internal/pkg/repository"
)

type StimusWordService interface {
	ListStimuses(ctx context.Context) ([]*datastruct.StimusWord, error)
}

type stimusWordService struct {
	dao repository.DAO
}

func (s *stimusWordService) ListStimuses(ctx context.Context) ([]*datastruct.StimusWord, error) {
	resp, err := s.dao.StimusWordQuery().List(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewStimusWordService(dao repository.DAO) StimusWordService {
	return &stimusWordService{dao: dao}
}
