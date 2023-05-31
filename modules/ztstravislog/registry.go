package ztstravislog

import "gorm.io/gorm"

func ZtsTravisLogRegistry(db *gorm.DB) Service {
	ztsTravisLogRepository := NewRepository(db)
	ztsTravisLogService := NewService(ztsTravisLogRepository)

	return ztsTravisLogService
}
