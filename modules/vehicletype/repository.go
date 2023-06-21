package vehicletype

import (
	"travis/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.VehicleType, error)
	Store(vehicleType domain.VehicleType) (domain.VehicleType, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.VehicleType, error) {
	var vehicleType []domain.VehicleType
	err := r.db.Find(&vehicleType).Error

	return vehicleType, err
}

func (r *repository) Store(vehicleType domain.VehicleType) (domain.VehicleType, error) {
	err := r.db.Create(&vehicleType).Error

	return vehicleType, err
}
