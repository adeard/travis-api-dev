package ztstravis

import (
	"travis/domain"

	"github.com/jinzhu/now"
)

type Service interface {
	GetAll(ztsTravisFilter domain.ZtsTravisFilter) ([]domain.ZtsTravis, error)
	Store(ztsTravisRequest domain.ZtsTravisRequest) (domain.ZtsTravis, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(ztsTravisFilter domain.ZtsTravisFilter) ([]domain.ZtsTravis, error) {

	if ztsTravisFilter.StartDate == "" {
		startDate := now.BeginningOfMonth()
		ztsTravisFilter.StartDate = startDate.Format("2006-01-02")
	}

	if ztsTravisFilter.EndDate == "" {
		endDate := now.EndOfMonth()
		ztsTravisFilter.EndDate = endDate.Format("2006-01-02")
	}

	ztsTravis, err := s.repository.FindAll(ztsTravisFilter)

	return ztsTravis, err
}

func (s *service) Store(ztsTravisRequest domain.ZtsTravisRequest) (domain.ZtsTravis, error) {

	ztsTravis, err := s.repository.Store(domain.ZtsTravis(ztsTravisRequest))

	return ztsTravis, err
}
