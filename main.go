package main

import (
	foodDetailsController "food-details-integrator-be/foodDetails/controller"
	"food-details-integrator-be/foodDetails/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_URL"),
		DialTimeout: time.Millisecond * 20,
	})

	r := gin.Default()

	cs := service.NewCacheService(*client)
	fds := service.NewFoodDetailsService(cs)
	detailsController := foodDetailsController.New(fds)

	r.GET("/food/:barcode", detailsController.GetFoodDetails)
	r.GET("/food/:barcode/kcal", detailsController.GetKcalForQuantity)

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "UP",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
