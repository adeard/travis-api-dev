package ztstravislog

import "travis/domain"

type Service interface {
	GetAll() ([]domain.ZtsTravisLog, error)
	Store(ztsTravisLogRequest domain.ZtsTravisLogRequest) (domain.ZtsTravisLog, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.ZtsTravisLog, error) {

	ztsTravisLog, err := s.repository.FindAll()

	return ztsTravisLog, err
}

func (s *service) Store(ztsTravisLogRequest domain.ZtsTravisLogRequest) (domain.ZtsTravisLog, error) {

	ztsTravisLog, err := s.repository.Store(domain.ZtsTravisLog(ztsTravisLogRequest))

	return ztsTravisLog, err
}
