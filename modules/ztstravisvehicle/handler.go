package ztstravisvehicle

import (
	"net/http"
	"travis/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ztsTravisVehicleHandler struct {
	ztsTravisVehicleService Service
}

func NewZtsTravisVehicleHandler(v1 *gin.RouterGroup, ztsTravisVehicleService Service) {

	handler := &ztsTravisVehicleHandler{ztsTravisVehicleService}

	ztsTravisVehicle := v1.Group("zts_travis/vehicle")

	ztsTravisVehicle.GET("", handler.GetAll)
	ztsTravisVehicle.POST("", handler.Store)
}

func (h *ztsTravisVehicleHandler) GetAll(c *gin.Context) {
	ztsTravisVehicle, err := h.ztsTravisVehicleService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravisVehicle,
	})
}

func (h *ztsTravisVehicleHandler) Store(c *gin.Context) {
	var ztsTravisVehicleRequest domain.ZtsTravisVehicleRequest

	c.ShouldBindJSON(&ztsTravisVehicleRequest)

	validate := validator.New()
	err := validate.Struct(ztsTravisVehicleRequest)
	if err != nil {
		errorMessages := []interface{}{}

		for _, v := range err.(validator.ValidationErrors) {
			errorArray := map[string]string{
				"field":   v.Field(),
				"message": v.ActualTag(),
			}

			errorMessages = append(errorMessages, errorArray)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	ztsTravisVehicle, err := h.ztsTravisVehicleService.Store(ztsTravisVehicleRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravisVehicle,
	})
}
