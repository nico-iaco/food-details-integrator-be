package service

import (
	"fmt"
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

func (s FoodDetailsService) GetProductDetails(barcode string) any {
	api := openfoodfacts.NewClient("world", "", "")
	//api.Sandbox()
	product, err := api.Product(barcode)
	if err != nil {
		fmt.Printf("error %s", err)
	}
	return product
}
