package ztstravis

import (
	"travis/domain"

	"github.com/jinzhu/now"
)

type Service interface {
	GetAll(ztsTravisFilter domain.ZtsTravisFilter) ([]domain.ZtsTravis, error)
	GetDetail(ztsTravisIdDetail domain.ZtsTravisIdDetail) (domain.ZtsTravis, error)
	Delete(ztsTravisIdDetail domain.ZtsTravisIdDetail) (domain.ZtsTravis, error)
	Store(ztsTravisRequest domain.ZtsTravisRequest) (domain.ZtsTravis, error)
	Update(ZtsTravisIdDetail domain.ZtsTravisIdDetail, ztsTravisUpdateInput domain.ZtsTravisUpdateRequest) (domain.ZtsTravis, error)
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

func (s *service) Update(ztsTravisIdDetail domain.ZtsTravisIdDetail, ztsTravisUpdateInput domain.ZtsTravisUpdateRequest) (domain.ZtsTravis, error) {

	ztsTravis, err := s.repository.FindDetail(ztsTravisIdDetail)
	if err != nil {
		return domain.ZtsTravis{}, err
	}

	if ztsTravisUpdateInput.NOTES_FR_TRANSP != "" {
		ztsTravis.NOTES_FR_TRANSP = ztsTravisUpdateInput.NOTES_FR_TRANSP
	}

	if ztsTravisUpdateInput.NOTES_FR_SALES != "" {
		ztsTravis.NOTES_FR_SALES = ztsTravisUpdateInput.NOTES_FR_SALES
	}

	newZtsTravis, err := s.repository.Update(ztsTravis, ztsTravisIdDetail)

	return newZtsTravis, err
}

func (s *service) GetDetail(ztsTravisIdDetail domain.ZtsTravisIdDetail) (domain.ZtsTravis, error) {

	ztsTravis, err := s.repository.FindDetail(ztsTravisIdDetail)
	if err != nil {
		return domain.ZtsTravis{}, err
	}

	return ztsTravis, err
}

func (s *service) Delete(ztsTravisIdDetail domain.ZtsTravisIdDetail) (domain.ZtsTravis, error) {

	ztsTravis, err := s.repository.FindDetail(ztsTravisIdDetail)
	if err != nil {
		return domain.ZtsTravis{}, err
	}

	_, err = s.repository.Delete(ztsTravisIdDetail)
	if err != nil {
		return domain.ZtsTravis{}, err
	}

	return ztsTravis, err
}
