package controller

import (
	foodDetailsService "food-details-integrator-be/foodDetails/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type foodDetailsController struct {
	foodDetailsService foodDetailsService.FoodDetailsService
}

var fdc *foodDetailsController = nil

func New(fds foodDetailsService.FoodDetailsService) foodDetailsController {
	if fdc == nil {
		fdc = new(foodDetailsController)
		fdc.foodDetailsService = fds
	}
	return *fdc
}

func (f foodDetailsController) GetFoodDetails(c *gin.Context) {
	barcode := c.Param("barcode")

	result := f.foodDetailsService.GetProductDetails(barcode)

	c.JSON(http.StatusOK, result)
}
