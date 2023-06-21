package ztstravisvehicle

import (
	"travis/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.ZtsTravisVehicle, error)
	Store(ztstravisvehicle domain.ZtsTravisVehicle) (domain.ZtsTravisVehicle, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.ZtsTravisVehicle, error) {
	var ztsTravisVehicle []domain.ZtsTravisVehicle
	err := r.db.Find(&ztsTravisVehicle).Error

	return ztsTravisVehicle, err
}

func (r *repository) Store(ztsTravisVehicle domain.ZtsTravisVehicle) (domain.ZtsTravisVehicle, error) {
	err := r.db.Create(&ztsTravisVehicle).Error

	return ztsTravisVehicle, err
}
