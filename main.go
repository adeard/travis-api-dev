package main

import (
	"travis/config"
	"travis/middlewares"
	"travis/modules/auth"
	"travis/modules/ztstravis"
	"travis/modules/ztstravisatta"
	"travis/modules/ztstravislog"

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
	ztstravisatta.NewZtsTravisAttaHandler(v1, ztstravisatta.ZtsTravisAttaRegistry(db))
	ztstravislog.NewZtsTravisLogHandler(v1, ztstravislog.ZtsTravisLogRegistry(db))

	router.Run(":86")
}
