package ztstravisdriver

import (
	"net/http"
	"travis/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ztsTravisDriverHandler struct {
	ztsTravisDriverService Service
}

func NewZtsTravisDriverHandler(v1 *gin.RouterGroup, ztsTravisDriverService Service) {

	handler := &ztsTravisDriverHandler{ztsTravisDriverService}

	ztsTravisDriver := v1.Group("zts_travis/driver")

	ztsTravisDriver.GET("", handler.GetAll)
	ztsTravisDriver.POST("", handler.Store)
}

func (h *ztsTravisDriverHandler) GetAll(c *gin.Context) {
	ztsTravisDriver, err := h.ztsTravisDriverService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravisDriver,
	})
}

func (h *ztsTravisDriverHandler) Store(c *gin.Context) {
	var ztsTravisDriverRequest domain.ZtsTravisDriverRequest

	c.ShouldBindJSON(&ztsTravisDriverRequest)

	validate := validator.New()
	err := validate.Struct(ztsTravisDriverRequest)
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

	ztsTravisDriver, err := h.ztsTravisDriverService.Store(ztsTravisDriverRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravisDriver,
	})
}
