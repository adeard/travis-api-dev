package ztstravis

import (
	"fmt"
	"net/http"
	"travis/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ztsTravisHandler struct {
	ztsTravisService Service
}

func NewZtsTravisHandler(v1 *gin.RouterGroup, ztsTravisService Service) {

	handler := &ztsTravisHandler{ztsTravisService}

	ztsTravis := v1.Group("zts_travis")

	ztsTravis.GET("", handler.GetAll)
	ztsTravis.POST("", handler.Store)
}

func (h *ztsTravisHandler) GetAll(c *gin.Context) {
	var input domain.ZtsTravisFilter

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ztsTravis, err := h.ztsTravisService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravis,
	})
}

func (h *ztsTravisHandler) Store(c *gin.Context) {
	var ztsTravisRequest domain.ZtsTravisRequest

	err := c.ShouldBindJSON(&ztsTravisRequest)
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

	ztsTravis, err := h.ztsTravisService.Store(ztsTravisRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravis,
	})
}
