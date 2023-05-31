package ztstravislog

import (
	"travis/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]domain.ZtsTravisLog, error)
	Store(ztsTravisLog domain.ZtsTravisLog) (domain.ZtsTravisLog, error)
	Update(ztsTravisLog domain.ZtsTravisLog) (domain.ZtsTravisLog, error)
	Delete(ztsTravisLog domain.ZtsTravisLog) (domain.ZtsTravisLog, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.ZtsTravisLog, error) {
	var ztsTravisLog []domain.ZtsTravisLog
	err := r.db.Find(&ztsTravisLog).Error

	return ztsTravisLog, err
}

func (r *repository) Store(ztsTravisLog domain.ZtsTravisLog) (domain.ZtsTravisLog, error) {
	err := r.db.Create(&ztsTravisLog).Error

	return ztsTravisLog, err
}

func (r *repository) Update(ztsTravisLog domain.ZtsTravisLog) (domain.ZtsTravisLog, error) {
	err := r.db.Save(&ztsTravisLog).Error

	return ztsTravisLog, err
}

func (r *repository) Delete(ztsTravisLog domain.ZtsTravisLog) (domain.ZtsTravisLog, error) {
	err := r.db.Delete(ztsTravisLog).Error

	return ztsTravisLog, err
}
