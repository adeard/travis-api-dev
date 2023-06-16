package ztstravis

import (
	"travis/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(ztsTravisFilter domain.ZtsTravisFilter) ([]domain.ZtsTravis, error)
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

func (r *repository) FindAll(ztsTravisFilter domain.ZtsTravisFilter) ([]domain.ZtsTravis, error) {
	var ztsTravis []domain.ZtsTravis

	q := r.db.Debug()
	q = q.Where("BLDAT between ? and ?", ztsTravisFilter.StartDate, ztsTravisFilter.EndDate)

	err := q.Find(&ztsTravis).Error

	return ztsTravis, err
}

func (r *repository) Store(ztsTravis domain.ZtsTravis) (domain.ZtsTravis, error) {
	err := r.db.Table("ZTS_TRAVIS").Debug().Create(&ztsTravis).Error

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
