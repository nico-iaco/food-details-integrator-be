package controller

import (
	foodDetailsService "food-details-integrator-be/foodDetails/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	result, err := f.foodDetailsService.GetProductDetails(barcode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, result)
}

func (f foodDetailsController) GetKcalForQuantity(c *gin.Context) {
	barcode := c.Param("barcode")
	quantity := c.Query("quantity")
	unit := c.Query("unit")

	quantityConverted, errParse := strconv.ParseFloat(quantity, 64)

	if errParse != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errParse.Error(),
		})
	}
	result, err := f.foodDetailsService.GetKcalForFoodQuantity(barcode, quantityConverted, unit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, result)
}
