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

func (s FoodDetailsService) GetProductDetails(barcode string) (model.FoodDetails, error) {
	api := openfoodfacts.NewClient("world", "", "")
	//api.Sandbox()
	product, err := api.Product(barcode)
	foodDetails := model.FoodDetails{}
	if err != nil {
		fmt.Printf("error %s \n", err)
		return model.FoodDetails{}, err
	} else {
		mappedField := smapping.MapTags(product, "json")
		err := smapping.FillStructByTags(&foodDetails, mappedField, "json")
		if err != nil {
			return model.FoodDetails{}, err
		}
		foodDetails.ImageURL = product.ImageURL.URL.String()
		foodDetails.ImageNutritionURL = product.ImageNutritionURL.URL.String()

	}
	return foodDetails, nil
}

func (s FoodDetailsService) GetKcalForFoodQuantity(barcode string, quantity float64, unit string) (float64, error) {
	foodDetails, err := s.GetProductDetails(barcode)

	if err != nil {
		return 0, err
	}
	kcalFor100 := foodDetails.Nutriments.EnergyKcal
	// quantity : x = 100 : kcalFor100
	computedQuantity := (kcalFor100 / 100) * quantity

	return computedQuantity, nil
}
