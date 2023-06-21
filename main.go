package main

import (
	"travis/config"
	"travis/middlewares"
	"travis/modules/auth"
	"travis/modules/vehicletype"
	"travis/modules/ztstravis"
	"travis/modules/ztstravisatta"
	"travis/modules/ztstravisdriver"
	"travis/modules/ztstravislog"
	"travis/modules/ztstravisvehicle"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db := config.Connect()

	router := gin.Default()
	router.Use(cors.AllowAll())

	v1 := router.Group("api/v1")
	auth.NewAuthHandler(v1, auth.AuthRegistry(db))

	v1.Use(middlewares.AuthService_Sample())

	ztstravis.NewZtsTravisHandler(v1, ztstravis.ZtsTravisRegistry(db))
	vehicletype.NewVehicleTypeHandler(v1, vehicletype.VehicleTypeRegistry(db))
	ztstravislog.NewZtsTravisLogHandler(v1, ztstravislog.ZtsTravisLogRegistry(db))
	ztstravisatta.NewZtsTravisAttaHandler(v1, ztstravisatta.ZtsTravisAttaRegistry(db))
	ztstravisdriver.NewZtsTravisDriverHandler(v1, ztstravisdriver.ZtsTravisDriverRegistry(db))
	ztstravisvehicle.NewZtsTravisVehicleHandler(v1, ztstravisvehicle.ZtsTravisVehicleRegistry(db))

	router.Run(":86")
}
