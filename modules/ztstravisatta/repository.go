package ztstravisatta

import (
	"travis/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.ZtsTravisAtta, error)
	Store(ztsTravisAtta domain.ZtsTravisAtta) (domain.ZtsTravisAtta, error)
	Update(ztsTravisAtta domain.ZtsTravisAtta) (domain.ZtsTravisAtta, error)
	Delete(ztsTravisAtta domain.ZtsTravisAtta) (domain.ZtsTravisAtta, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.ZtsTravisAtta, error) {
	var ztsTravisAtta []domain.ZtsTravisAtta
	err := r.db.Find(&ztsTravisAtta).Error

	return ztsTravisAtta, err
}

func (r *repository) Store(ztsTravisAtta domain.ZtsTravisAtta) (domain.ZtsTravisAtta, error) {
	err := r.db.Create(&ztsTravisAtta).Error

	return ztsTravisAtta, err
}

func (r *repository) Update(ztsTravisAtta domain.ZtsTravisAtta) (domain.ZtsTravisAtta, error) {
	err := r.db.Save(&ztsTravisAtta).Error

	return ztsTravisAtta, err
}

func (r *repository) Delete(ztsTravisAtta domain.ZtsTravisAtta) (domain.ZtsTravisAtta, error) {
	err := r.db.Delete(ztsTravisAtta).Error

	return ztsTravisAtta, err
}
