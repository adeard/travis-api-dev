package ztstravisvehicle

import "gorm.io/gorm"

func ZtsTravisVehicleRegistry(db *gorm.DB) Service {
	ztsTravisVehicleRepository := NewRepository(db)
	ztsTravisVehicleService := NewService(ztsTravisVehicleRepository)

	return ztsTravisVehicleService
}
