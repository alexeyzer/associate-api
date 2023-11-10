package service

import (
	"github.com/alexeyzer/associate-api/internal/pkg/repository"
)

type AssociateWordService interface {
}

type  associateWordService struct {
	dao repository.DAO
}


func NewAssociateWordService(dao repository.DAO) AssociateWordService {
	return &associateWordService{dao: dao}
}
