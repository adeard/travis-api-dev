package ztstravisdriver

import "travis/domain"

type Service interface {
	GetAll() ([]domain.ZtsTravisDriver, error)
	Store(ztsTravisDriverRequest domain.ZtsTravisDriverRequest) (domain.ZtsTravisDriver, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.ZtsTravisDriver, error) {

	ztsTravisDriver, err := s.repository.FindAll()

	return ztsTravisDriver, err
}

func (s *service) Store(ztsTravisDriverRequest domain.ZtsTravisDriverRequest) (domain.ZtsTravisDriver, error) {

	ztsTravisDriver, err := s.repository.Store(domain.ZtsTravisDriver(ztsTravisDriverRequest))

	return ztsTravisDriver, err
}
