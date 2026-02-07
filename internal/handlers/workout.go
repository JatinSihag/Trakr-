package handlers

import (
	"net/http"

	"github.com/JatinSihag/Trakr/internal/db"
	"github.com/JatinSihag/Trakr/internal/models"
	"github.com/gin-gonic/gin"
)

func LogWorkout(c *gin.Context) {
	userId, _ := c.Get("userId")
	if userId == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "No token found"})
		return
	}
	var input models.Workouts
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid payload"})
		return
	}
	totalBurned := 0

	for _, s := range input.Sets {
		totalBurned += s.Calories
	}

	tx, err := db.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to connect DB"})
		return
	}
	queryHeader := `INSERT into workouts (user_id,workout_name,start_time,end_time,notes,calories_burned)
	VALUES (?,?,?,?,?,?)`
	result, err := tx.Exec(queryHeader, userId, input.WorkoutName, input.StartTime, input.EndTime, input.Notes, totalBurned)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to save workout"})
		return
	}
	workoutId, _ := result.LastInsertId()
	querySet := `INSERT into workout_sets(workout_id,exercise_id,set_number,weight,reps,is_pr,duration_minutes,set_calories) VALUES(?,?,?,?,?,?,?,?)`
	for i, s := range input.Sets {
		_, err := tx.Exec(querySet, workoutId, s.ExerciseId, i+1, s.Weight, s.Reps, s.IsPr, s.Duration, s.Calories)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to Save Set"})
			return
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"message":        "Workout Saved",
		"total_calories": totalBurned,
	})
}
