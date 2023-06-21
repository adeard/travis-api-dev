package ztstravisdriver

import (
	"travis/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.ZtsTravisDriver, error)
	Store(ztsTravisDriver domain.ZtsTravisDriver) (domain.ZtsTravisDriver, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.ZtsTravisDriver, error) {
	var ztsTravisDriver []domain.ZtsTravisDriver
	err := r.db.Find(&ztsTravisDriver).Error

	return ztsTravisDriver, err
}

func (r *repository) Store(ztsTravisDriver domain.ZtsTravisDriver) (domain.ZtsTravisDriver, error) {
	err := r.db.Create(&ztsTravisDriver).Error

	return ztsTravisDriver, err
}
