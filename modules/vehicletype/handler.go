package vehicletype

import (
	"fmt"
	"net/http"
	"travis/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type vehicleTypeHandler struct {
	vehicleTypeService Service
}

func NewVehicleTypeHandler(v1 *gin.RouterGroup, vehicleTypeService Service) {

	handler := &vehicleTypeHandler{vehicleTypeService}

	vehicleType := v1.Group("vehicle_type")

	vehicleType.GET("", handler.GetAll)
	vehicleType.POST("", handler.Store)
}

func (h *vehicleTypeHandler) GetAll(c *gin.Context) {
	vehicleType, err := h.vehicleTypeService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": vehicleType,
	})
}

func (h *vehicleTypeHandler) Store(c *gin.Context) {
	var vehicleTypeRequest domain.VehicleTypeRequest

	c.ShouldBindJSON(&vehicleTypeRequest)

	fmt.Println(vehicleTypeRequest)

	validate := validator.New()
	err := validate.Struct(vehicleTypeRequest)
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

	vehicleType, err := h.vehicleTypeService.Store(vehicleTypeRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": vehicleType,
	})
}
