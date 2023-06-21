package ztstravisvehicle

import "travis/domain"

type Service interface {
	GetAll() ([]domain.ZtsTravisVehicle, error)
	Store(ztsTravisVehicleRequest domain.ZtsTravisVehicleRequest) (domain.ZtsTravisVehicle, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.ZtsTravisVehicle, error) {

	ztsTravisVehicle, err := s.repository.FindAll()

	return ztsTravisVehicle, err
}

func (s *service) Store(ztsTravisVehicleRequest domain.ZtsTravisVehicleRequest) (domain.ZtsTravisVehicle, error) {

	ztsTravisVehicle, err := s.repository.Store(domain.ZtsTravisVehicle(ztsTravisVehicleRequest))

	return ztsTravisVehicle, err
}
