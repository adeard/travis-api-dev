package ztstravisatta

import (
	"fmt"
	"net/http"
	"travis/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ztsTravisAttaHandler struct {
	ztsTravisAttaService Service
}

func NewZtsTravisAttaHandler(v1 *gin.RouterGroup, ztsTravisAttaService Service) {

	handler := &ztsTravisAttaHandler{ztsTravisAttaService}

	ztsTravisAtta := v1.Group("zts_travis/atta")

	ztsTravisAtta.GET("", handler.GetAll)
	ztsTravisAtta.POST("", handler.Store)
}

func (h *ztsTravisAttaHandler) GetAll(c *gin.Context) {
	ztsTravis, err := h.ztsTravisAttaService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravis,
	})
}

func (h *ztsTravisAttaHandler) Store(c *gin.Context) {
	var ztsTravisAttaRequest domain.ZtsTravisAttaRequest

	err := c.ShouldBindJSON(&ztsTravisAttaRequest)
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

	ztsTravisAtta, err := h.ztsTravisAttaService.Store(ztsTravisAttaRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravisAtta,
	})
}
