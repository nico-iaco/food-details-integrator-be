package service

import (
	"fmt"
	"food-details-integrator-be/foodDetails/model"
	"github.com/mashingan/smapping"
	"github.com/openfoodfacts/openfoodfacts-go"
)

type FoodDetailsService struct {
}

var fds *FoodDetailsService = nil

func New() FoodDetailsService {
	if fds == nil {
		fds = new(FoodDetailsService)
	}
	return *fds
}

func (s FoodDetailsService) GetProductDetails(barcode string) model.FoodDetails {
	api := openfoodfacts.NewClient("world", "", "")
	//api.Sandbox()
	product, err := api.Product(barcode)
	foodDetails := model.FoodDetails{}
	if err != nil {
		fmt.Printf("error %s", err)
	} else {
		mappedField := smapping.MapTags(product, "json")

		err := smapping.FillStructByTags(&foodDetails, mappedField, "json")
		foodDetails.ImageURL = product.ImageURL.URL.String()
		foodDetails.ImageNutritionURL = product.ImageNutritionURL.URL.String()
		if err != nil {
			return model.FoodDetails{}
		}
	}
	return foodDetails
}
