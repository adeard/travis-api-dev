package ztstravislog

import (
	"fmt"
	"net/http"
	"travis/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ztsTravisLogHandler struct {
	ztsTravisLogService Service
}

func NewZtsTravisLogHandler(v1 *gin.RouterGroup, ztsTravisLogService Service) {

	handler := &ztsTravisLogHandler{ztsTravisLogService}

	ztsTravisLog := v1.Group("zts_travis/log")

	ztsTravisLog.GET("", handler.GetAll)
	ztsTravisLog.POST("", handler.Store)
}

func (h *ztsTravisLogHandler) GetAll(c *gin.Context) {
	ztsTravis, err := h.ztsTravisLogService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravis,
	})
}

func (h *ztsTravisLogHandler) Store(c *gin.Context) {
	var ztsTravisLogRequest domain.ZtsTravisLogRequest

	err := c.ShouldBindJSON(&ztsTravisLogRequest)
	if err != nil {

		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s , condition : %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": errorMessages,
		})
	}

	ztsTravisLog, err := h.ztsTravisLogService.Store(ztsTravisLogRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravisLog,
	})
}
