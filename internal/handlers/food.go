package handlers

import (
	"database/sql"
	"net/http"

	"github.com/JatinSihag/Trakr/internal/db"
	"github.com/JatinSihag/Trakr/internal/models"
	"github.com/gin-gonic/gin"
)

func SearchFood(c *gin.Context) {
	queryParam := c.Query("query")
	if queryParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Query parameter"})
		return
	}
	query := `SELECT food_id,food_name,calories_per_100g,protein_per_100g,fats_per_100g,carbs_per_100g FROM foods
	WHERE food_name LIKE ? LIMIT 10`

	rows, err := db.DB.Query(query, "%"+queryParam+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
		return
	}
	defer rows.Close()
	var results []models.Food
	for rows.Next() {
		var f models.Food
		err := rows.Scan(&f.FoodId, &f.FoodName, &f.Calories, &f.Protein, &f.Fats, &f.Carbs)
		if err != nil {
			continue
		}
		results = append(results, f)
	}
	c.JSON(http.StatusOK, results)
}

func LogFood(c *gin.Context) {
	userId, _ := c.Get("userId")
	if userId == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		return
	}
	var input models.FoodLog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Input"})
		return
	}
	validMeals := map[string]bool{"breakfast": true, "lunch": true, "dinner": true, "snacks": true}
	if !validMeals[input.MealType] {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid Meal": "Enter Appropriate meal"})
		return
	}
	query := `INSERT INTO nutrition_logs (user_id,food_id,quantity,log_date,meal_type)
	VALUES(?,?,?,CURRENT_DATE,?)`
	_, err := db.DB.Exec(query, userId, input.FoodId, input.Quantity, input.MealType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to log food"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success": "Food Logged Successfully"})
}

func GetDailyLog(c *gin.Context) {
	userId, _ := c.Get("userId")
	date := c.Query("date")

	var rows *sql.Rows
	var err error

	query := `SELECT l.log_id, l.food_id, l.quantity, l.meal_type, l.log_date, 
                     f.food_name, f.calories_per_100g
              FROM nutrition_logs l
              JOIN foods f ON l.food_id = f.food_id
              WHERE l.user_id = ? AND l.log_date = `

	if date == "" {
		query += "CURRENT_DATE"
		rows, err = db.DB.Query(query, userId)
	} else {
		query += "?"
		rows, err = db.DB.Query(query, userId, date)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get foodLogs"})
		return
	}
	defer rows.Close()

	var logs []models.FoodLog
	var totalCalories float64 = 0

	for rows.Next() {
		var l models.FoodLog
		var calPer100g float64

		err := rows.Scan(&l.LogId, &l.FoodId, &l.Quantity, &l.MealType, &l.LogDate, &l.FoodName, &calPer100g)
		if err != nil {
			continue
		}
		l.TotalCalories = (calPer100g * l.Quantity) / 100
		totalCalories += l.TotalCalories

		logs = append(logs, l)
	}
	if date == "" {
		date = "Today"
	}

	c.JSON(http.StatusOK, gin.H{
		"date":          date,
		"totalCalories": totalCalories,
		"logs":          logs,
	})
}
