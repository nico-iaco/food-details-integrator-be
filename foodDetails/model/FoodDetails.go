package model

type FoodDetails struct {
	GenericName       string         `json:"generic_name"`
	ImageURL          string         `json:"image_url"`
	ImageNutritionURL string         `json:"image_nutrition_url"`
	Nutriments        FoodNutriments `json:"nutriments"`
	Quantity          string         `json:"quantity"`
}
