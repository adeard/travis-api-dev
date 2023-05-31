package ztstravis

import (
	"travis/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.ZtsTravis, error)
	Store(ztsTravis domain.ZtsTravis) (domain.ZtsTravis, error)
	Update(ztsTravis domain.ZtsTravis) (domain.ZtsTravis, error)
	Delete(ztsTravis domain.ZtsTravis) (domain.ZtsTravis, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.ZtsTravis, error) {
	var ztsTravis []domain.ZtsTravis
	err := r.db.Find(&ztsTravis).Error

	return ztsTravis, err
}

func (r *repository) Store(ztsTravis domain.ZtsTravis) (domain.ZtsTravis, error) {
	err := r.db.Create(&ztsTravis).Error

	return ztsTravis, err
}

func (r *repository) Update(ztsTravis domain.ZtsTravis) (domain.ZtsTravis, error) {
	err := r.db.Save(&ztsTravis).Error

	return ztsTravis, err
}

func (r *repository) Delete(ztsTravis domain.ZtsTravis) (domain.ZtsTravis, error) {
	err := r.db.Delete(ztsTravis).Error

	return ztsTravis, err
}
