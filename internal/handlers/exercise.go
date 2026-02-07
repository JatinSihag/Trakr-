package handlers

import (
	"database/sql"
	"net/http"

	"github.com/JatinSihag/Trakr/internal/db"
	"github.com/JatinSihag/Trakr/internal/models"
	"github.com/gin-gonic/gin"
)

func GetExercise(c *gin.Context) {
	userId, _ := c.Get("userId")
	if userId == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "No token found"})
		return
	}
	queryParam := c.Query("q")

	var rows *sql.Rows
	var err error
	if queryParam == "" {
		query := `SELECT exercise_id,exercise_name,body_part,category FROM exercises WHERE user_id IS NULL OR user_id=? LIMIT 20`
		rows, err = db.DB.Query(query, userId)
	} else {
		searchItem := "%" + queryParam + "%"
		query := `SELECT exercise_id, exercise_name, body_part, category 
                  FROM exercises 
                  WHERE (user_id IS NULL OR user_id=?) 
                  AND exercise_name LIKE ?`
		rows, err = db.DB.Query(query, userId, searchItem)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database Error"})
		return
	}
	defer rows.Close()
	var result []models.Exercise
	for rows.Next() {
		var e models.Exercise
		if err := rows.Scan(&e.ExerciseId, &e.ExerciseName, &e.BodyPart, &e.Category); err != nil {
			continue
		}
		result = append(result, e)
	}
	if result == nil {
		result = []models.Exercise{}
	}
	c.JSON(http.StatusOK, result)
}
