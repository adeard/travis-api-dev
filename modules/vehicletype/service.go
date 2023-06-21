package vehicletype

import "travis/domain"

type Service interface {
	GetAll() ([]domain.VehicleType, error)
	Store(vehicleTypeRequest domain.VehicleTypeRequest) (domain.VehicleType, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.VehicleType, error) {

	vehicleType, err := s.repository.FindAll()

	return vehicleType, err
}

func (s *service) Store(vehicleTypeRequest domain.VehicleTypeRequest) (domain.VehicleType, error) {

	vehicleType, err := s.repository.Store(domain.VehicleType(vehicleTypeRequest))

	return vehicleType, err
}
