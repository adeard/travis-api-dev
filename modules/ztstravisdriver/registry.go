package ztstravisdriver

import "gorm.io/gorm"

func ZtsTravisDriverRegistry(db *gorm.DB) Service {
	ztsTravisDriverRepository := NewRepository(db)
	ztsTravisDriverService := NewService(ztsTravisDriverRepository)

	return ztsTravisDriverService
}
