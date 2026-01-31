package handlers

import (
	"net/http"

	"github.com/JatinSihag/Trakr/internal/db"
	"github.com/JatinSihag/Trakr/internal/models"
	"github.com/gin-gonic/gin"
)

func UpdateProfile(c *gin.Context) {
	userId,exists :=c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized,gin.H{"message":"Unauthorized"})
		return
	}

	var input models.User
	// how shouldbind json works?
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Error":"Invalid Input"})
		return
	}
	query:= `UPDATE users SET weight=?,height=?,activity_level=?,goal=?,age=?,gender=? WHERE user_id = ?`
	_,err := db.DB.Exec(query,input.Weight,input.Height,input.ActivityLevel,input.Goal,input.Age,input.Gender,userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"Error":"Failed to update db"})
		return
	}

	userCalc:= models.User{
		Weight: input.Weight,
		Height: input.Height,
		ActivityLevel: input.ActivityLevel,
		Goal: input.Goal,
		Age: input.Age,
		Gender: input.Gender,
	}

	dailyCalories := userCalc.CalculateTDEE()

	c.JSON(http.StatusOK,gin.H{
		"Message":"User Profile Updated",
		"Calories Per day ":dailyCalories,
	})

}