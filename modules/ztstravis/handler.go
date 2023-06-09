package ztstravis

import (
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
	ztsTravis.GET(":mandt/:taskid/:vkorg/:werks/:vbeln/:posnr", handler.GetDetail)
	ztsTravis.PUT(":mandt/:taskid/:vkorg/:werks/:vbeln/:posnr", handler.Update)
	ztsTravis.DELETE(":mandt/:taskid/:vkorg/:werks/:vbeln/:posnr", handler.Delete)
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

	c.ShouldBindJSON(&ztsTravisRequest)

	validate := validator.New()
	err := validate.Struct(ztsTravisRequest)
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

	ztsTravis, err := h.ztsTravisService.Store(ztsTravisRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravis,
	})
}

func (h *ztsTravisHandler) Update(c *gin.Context) {
	var ztsTravisUpdateInput domain.ZtsTravisUpdateRequest

	c.ShouldBind(&ztsTravisUpdateInput)

	validate := validator.New()
	err := validate.Struct(ztsTravisUpdateInput)
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

	ztsTravisIdDetail := domain.ZtsTravisIdDetail{
		MANDT:  c.Param("mandt"),
		TASKID: c.Param("taskid"),
		VKORG:  c.Param("vkorg"),
		WERKS:  c.Param("werks"),
		VBELN:  c.Param("vbeln"),
		POSNR:  c.Param("posnr"),
	}

	result, err := h.ztsTravisService.Update(ztsTravisIdDetail, ztsTravisUpdateInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *ztsTravisHandler) GetDetail(c *gin.Context) {
	ztsTravisIdDetail := domain.ZtsTravisIdDetail{
		MANDT:  c.Param("mandt"),
		TASKID: c.Param("taskid"),
		VKORG:  c.Param("vkorg"),
		WERKS:  c.Param("werks"),
		VBELN:  c.Param("vbeln"),
		POSNR:  c.Param("posnr"),
	}

	result, err := h.ztsTravisService.GetDetail(ztsTravisIdDetail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *ztsTravisHandler) Delete(c *gin.Context) {
	ztsTravisIdDetail := domain.ZtsTravisIdDetail{
		MANDT:  c.Param("mandt"),
		TASKID: c.Param("taskid"),
		VKORG:  c.Param("vkorg"),
		WERKS:  c.Param("werks"),
		VBELN:  c.Param("vbeln"),
		POSNR:  c.Param("posnr"),
	}
	result, err := h.ztsTravisService.Delete(ztsTravisIdDetail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"data": result,
	})
}
