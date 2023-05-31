package ztstravisatta

import "travis/domain"

type Service interface {
	GetAll() ([]domain.ZtsTravisAtta, error)
	Store(ztsTravisAttaRequest domain.ZtsTravisAttaRequest) (domain.ZtsTravisAtta, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.ZtsTravisAtta, error) {

	ztsTravisAtta, err := s.repository.FindAll()

	return ztsTravisAtta, err
}

func (s *service) Store(ztsTravisAttaRequest domain.ZtsTravisAttaRequest) (domain.ZtsTravisAtta, error) {

	ztsTravisAtta, err := s.repository.Store(domain.ZtsTravisAtta(ztsTravisAttaRequest))

	return ztsTravisAtta, err
}
