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

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
