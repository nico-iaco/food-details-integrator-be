package model

type FoodNutriments struct {
	Salt                   float64 `json:"salt"`
	SaltUnit               string  `json:"salt_unit"`
	Sugars                 float64 `json:"sugars"`
	SugarsUnit             string  `json:"sugars_unit"`
	Iron                   float64 `json:"iron"`
	IronLabel              string  `json:"iron_label"`
	IronUnit               string  `json:"iron_unit"`
	CalciumUnit            string  `json:"calcium_unit"`
	Calcium                float64 `json:"calcium"`
	CalciumLabel           string  `json:"calcium_label"`
	Cholesterol100G        float64 `json:"cholesterol_100g"`
	SaturatedFat           float64 `json:"saturated-fat"`
	SaturatedFatUnit       string  `json:"saturated-fat_unit"`
	FatUnit                string  `json:"fat_unit"`
	Fat                    float64 `json:"fat"`
	TransFatLabel          string  `json:"trans-fat_label"`
	TransFatUnit           string  `json:"trans-fat_unit"`
	TransFat               float64 `json:"trans-fat"`
	VitaminA               float64 `json:"vitamin-a"`
	VitaminAUnit           string  `json:"vitamin-a_unit"`
	VitaminALabel          string  `json:"vitamin-a_label"`
	VitaminCUnit           string  `json:"vitamin-c_unit"`
	VitaminC               float64 `json:"vitamin-c"`
	VitaminCLabel          string  `json:"vitamin-c_label"`
	ProteinsUnit           string  `json:"proteins_unit"`
	Proteins               float64 `json:"proteins"`
	PolyunsaturatedFat100G float64 `json:"polyunsaturated-fat_100g"`
	Potassium100G          float64 `json:"potassium_100g"`
	Sodium                 float64 `json:"sodium"`
	SodiumUnit             string  `json:"sodium_unit"`
	CarbohydratesUnit      string  `json:"carbohydrates_unit"`
	Carbohydrates          float64 `json:"carbohydrates"`
	AlcoholUnit            string  `json:"alcohol_unit"`
	Alcohol                float64 `json:"alcohol"`
	MonounsaturatedFat100G float64 `json:"monounsaturated-fat_100g"`
	NovaGroup              float64 `json:"nova-group"`
	EnergyKcal             float64 `json:"energy-kcal"`
	EnergyKcalUnit         string  `json:"energy-kcal_unit"`
	NutritionScoreFr       float64 `json:"nutrition-score-fr"`
	NutritionScoreUk       float64 `json:"nutrition-score-uk"`
	Fiber                  float64 `json:"fiber"`
	FiberUnit              string  `json:"fiber_unit"`
}
