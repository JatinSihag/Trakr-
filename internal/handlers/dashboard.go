package handlers

import (
	"database/sql"
	"net/http"
	"sync"
	"time"

	"github.com/JatinSihag/Trakr/internal/db"
	"github.com/JatinSihag/Trakr/internal/models"
	"github.com/gin-gonic/gin"
)

type DashboardData struct {
	Date              string  `json:"date"`
	DailyTarget       float64 `json:"daily_target"`
	CaloriesConsumed  float64 `json:"calories_consumed"`
	CaloriesBurned    float64 `json:"calories_burned"`
	CaloriesRemaining float64 `json:"calories_remaining"`
	Status            string  `json:"status"`
}

func GetDashboard(c *gin.Context) {
	userId, _ := c.Get("userId")
	if userId == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		return
	}

	var wg sync.WaitGroup
	wg.Add(3)

	userChan := make(chan *models.User, 1)
	caloriesChan := make(chan float64, 1)
	exerciseChan := make(chan float64, 1)
	errChan := make(chan error, 3)

	go func() {
		defer wg.Done()
		var u models.User
		query := `SELECT 
			COALESCE(weight, 0), 
			COALESCE(height, 0), 
			COALESCE(age, 0), 
			COALESCE(gender, 'unknown'), 
			COALESCE(activity_level, 'Not set'), 
			COALESCE(goal, 'Not set') 
			FROM users WHERE user_id = ?`

		err := db.DB.QueryRow(query, userId).Scan(&u.Weight, &u.Height, &u.Age, &u.Gender, &u.ActivityLevel, &u.Goal)
		if err != nil {
			errChan <- err
			return
		}
		userChan <- &u
	}()

	go func() {
		defer wg.Done()
		query := `SELECT SUM((f.calories_per_100g * l.quantity)/100)
        FROM food_logs l
        JOIN foods f ON l.food_id=f.food_id
        WHERE l.user_id = ? AND l.log_date = CURRENT_DATE`

		var totalConsumed sql.NullFloat64
		err := db.DB.QueryRow(query, userId).Scan(&totalConsumed)
		if err != nil && err != sql.ErrNoRows {
			errChan <- err
			return
		}
		if totalConsumed.Valid {
			caloriesChan <- totalConsumed.Float64
		} else {
			caloriesChan <- 0
		}
	}()

	go func() {
		defer wg.Done()
		query := `SELECT SUM(calories_burned) FROM workouts WHERE user_id = ? AND DATE(start_time)=CURRENT_DATE`
		var total sql.NullFloat64
		err := db.DB.QueryRow(query, userId).Scan(&total)
		if err != nil && err != sql.ErrNoRows {
			errChan <- err
			return
		}
		if total.Valid {
			exerciseChan <- total.Float64
		} else {
			exerciseChan <- 0
		}
	}()

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to load dashboard data"})
		return
	}

	user := <-userChan
	consumed := <-caloriesChan
	burned := <-exerciseChan

	tdee := user.CalculateTDEE()
	remaining := (tdee + burned) - consumed

	c.JSON(http.StatusOK, DashboardData{
		Date:              time.Now().Format("2006-01-02"),
		DailyTarget:       tdee,
		CaloriesConsumed:  consumed,
		CaloriesRemaining: remaining,
		CaloriesBurned:    burned,
		Status:            user.Goal,
	})
}
