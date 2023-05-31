package ztstravisatta

import "gorm.io/gorm"

func ZtsTravisAttaRegistry(db *gorm.DB) Service {
	ztsTravisAttaRepository := NewRepository(db)
	ztsTravisAttaService := NewService(ztsTravisAttaRepository)

	return ztsTravisAttaService
}
