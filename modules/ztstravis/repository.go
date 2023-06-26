package ztstravis

import (
	"travis/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(ztsTravisFilter domain.ZtsTravisFilter) ([]domain.ZtsTravis, error)
	FindDetail(ztsTravisIdDetail domain.ZtsTravisIdDetail) (domain.ZtsTravis, error)
	Store(ztsTravis domain.ZtsTravis) (domain.ZtsTravis, error)
	Update(ztsTravis domain.ZtsTravis, ztsTravisIdDetail domain.ZtsTravisIdDetail) (domain.ZtsTravis, error)
	Delete(ztsTravisIdDetail domain.ZtsTravisIdDetail) (domain.ZtsTravisIdDetail, error)
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

	if ztsTravisFilter.TaskStatus != "" {
		q = q.Where("TASK_STATUS = ?", ztsTravisFilter.TaskStatus)
	}

	err := q.Find(&ztsTravis).Error

	return ztsTravis, err
}

func (r *repository) Store(ztsTravis domain.ZtsTravis) (domain.ZtsTravis, error) {
	err := r.db.Table("ZTS_TRAVIS").Debug().Create(&ztsTravis).Error

	return ztsTravis, err
}

func (r *repository) Update(ztsTravis domain.ZtsTravis, ztsTravisIdDetail domain.ZtsTravisIdDetail) (domain.ZtsTravis, error) {
	q := r.db.Debug()

	if ztsTravisIdDetail.MANDT != "" {
		q = q.Where("MANDT = ?", ztsTravisIdDetail.MANDT)
	}

	if ztsTravisIdDetail.TASKID != "" {
		q = q.Where("TASKID = ?", ztsTravisIdDetail.TASKID)
	}

	if ztsTravisIdDetail.VKORG != "" {
		q = q.Where("VKORG = ?", ztsTravisIdDetail.VKORG)
	}

	if ztsTravisIdDetail.WERKS != "" {
		q = q.Where("WERKS = ?", ztsTravisIdDetail.WERKS)
	}

	if ztsTravisIdDetail.VBELN != "" {
		q = q.Where("VBELN = ?", ztsTravisIdDetail.VBELN)
	}

	if ztsTravisIdDetail.POSNR != "" {
		q = q.Where("POSNR = ?", ztsTravisIdDetail.POSNR)
	}

	err := q.Save(&ztsTravis).Error

	return ztsTravis, err
}

func (r *repository) Delete(ztsTravisIdDetail domain.ZtsTravisIdDetail) (domain.ZtsTravisIdDetail, error) {
	q := r.db.Debug().Table("ZTS_TRAVIS")

	if ztsTravisIdDetail.MANDT != "" {
		q = q.Where("MANDT = ?", ztsTravisIdDetail.MANDT)
	}

	if ztsTravisIdDetail.TASKID != "" {
		q = q.Where("TASKID = ?", ztsTravisIdDetail.TASKID)
	}

	if ztsTravisIdDetail.VKORG != "" {
		q = q.Where("VKORG = ?", ztsTravisIdDetail.VKORG)
	}

	if ztsTravisIdDetail.WERKS != "" {
		q = q.Where("WERKS = ?", ztsTravisIdDetail.WERKS)
	}

	if ztsTravisIdDetail.VBELN != "" {
		q = q.Where("VBELN = ?", ztsTravisIdDetail.VBELN)
	}

	if ztsTravisIdDetail.POSNR != "" {
		q = q.Where("POSNR = ?", ztsTravisIdDetail.POSNR)
	}

	err := q.Delete(ztsTravisIdDetail).Error

	return ztsTravisIdDetail, err
}

func (r *repository) FindDetail(ztsTravisIdDetail domain.ZtsTravisIdDetail) (domain.ZtsTravis, error) {
	var ztsTravis domain.ZtsTravis

	q := r.db.Debug().Table("ZTS_TRAVIS").
		Where("MANDT = ?", ztsTravisIdDetail.MANDT).
		Where("TASKID = ?", ztsTravisIdDetail.TASKID).
		Where("VKORG = ?", ztsTravisIdDetail.VKORG).
		Where("WERKS = ?", ztsTravisIdDetail.WERKS).
		Where("VBELN = ?", ztsTravisIdDetail.VBELN).
		Where("POSNR = ?", ztsTravisIdDetail.POSNR)

	err := q.First(&ztsTravis).Error

	return ztsTravis, err
}
