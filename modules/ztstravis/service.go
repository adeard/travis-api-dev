package ztstravis

import "travis/domain"

type Service interface {
	GetAll() ([]domain.ZtsTravis, error)
	Store(ztsTravisRequest domain.ZtsTravisRequest) (domain.ZtsTravis, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.ZtsTravis, error) {

	ztsTravis, err := s.repository.FindAll()

	return ztsTravis, err
}

func (s *service) Store(ztsTravisRequest domain.ZtsTravisRequest) (domain.ZtsTravis, error) {

	ztsTravis, err := s.repository.Store(domain.ZtsTravis(ztsTravisRequest))

	return ztsTravis, err
}
