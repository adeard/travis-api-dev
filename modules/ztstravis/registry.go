package ztstravis

import "gorm.io/gorm"

func ZtsTravisRegistry(db *gorm.DB) Service {
	ztsTravisRepository := NewRepository(db)
	ztsTravisService := NewService(ztsTravisRepository)

	return ztsTravisService
}
