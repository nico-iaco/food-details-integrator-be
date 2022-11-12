package service

import (
	"fmt"
	"food-details-integrator-be/foodDetails/model"
	"github.com/mashingan/smapping"
	"github.com/openfoodfacts/openfoodfacts-go"
	"os"
)

type FoodDetailsService struct {
	CacheService CacheService
}

var fds *FoodDetailsService = nil

func NewFoodDetailsService(cs CacheService) FoodDetailsService {
	if fds == nil {
		fds = new(FoodDetailsService)
		fds.CacheService = cs
	}
	return *fds
}

func (s FoodDetailsService) GetProductDetails(barcode string) (model.FoodDetails, error) {
	foodDetails := model.FoodDetails{}
	var err error

	isRedisEnabled := os.Getenv("REDIS_ENABLED") == "true"

	if isRedisEnabled {
		err = cs.GetObject(barcode, &foodDetails)
		if err != nil {
			foodDetails, err = s.getFoodDetailsFromApi(barcode)
		} else {
			fmt.Println("cache hit")
		}
		err = cs.SetObject(barcode, foodDetails)
	} else {
		foodDetails, err = s.getFoodDetailsFromApi(barcode)
	}
	return foodDetails, err
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

func (s FoodDetailsService) getFoodDetailsFromApi(barcode string) (model.FoodDetails, error) {
	api := openfoodfacts.NewClient("world", "", "")
	foodDetails := model.FoodDetails{}

	if os.Getenv("IS_SANDBOX") == "true" {
		api.Sandbox()
	}

	product, err := api.Product(barcode)

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
