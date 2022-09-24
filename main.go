package main

import (
	foodDetailsController "food-details-integrator-be/foodDetails/controller"
	foodDetailsService "food-details-integrator-be/foodDetails/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fds := foodDetailsService.New()
	detailsController := foodDetailsController.New(fds)

	r.GET("/food/:barcode", detailsController.GetFoodDetails)
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "UP",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
