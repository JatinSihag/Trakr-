package models

type Food struct {
	FoodId   int     `json:"food_id"`
	FoodName string  `json:"food_name"`
	Calories float64 `json:"calories_per_100g"`
	Fats     float64 `json:"fats_per_100g"`
	Protein  float64 `json:"protein_per_100g"`
	Carbs    float64 `json:"carbs_per_100g"`
}

type FoodLog struct {
	LogId     int  `json:"log_id"`
	UserId    int  `json:"user_id"`
	FoodId    int  `json:"food_id"`
	PortionId *int `json:"portion_id,omitempty"`

	Quantity float64 `json:"quantity"`  // decimal(4,1)
	LogDate  string  `json:"log_date"`  // YYYY-MM-DD
	MealType string  `json:"meal_type"` // ENUM: breakfast, lunch, dinner, snacks

	FoodName      string  `json:"food_name,omitempty"`
	TotalCalories float64 `json:"total_calories,omitempty"`
}
