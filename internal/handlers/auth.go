package handlers

import (
	"net/http"

	"github.com/JatinSihag/Trakr/internal/db"
	"github.com/JatinSihag/Trakr/internal/models"
	"github.com/JatinSihag/Trakr/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data input"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}
	err = insertIntoDb(user, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Success": "User registered successfully"})
}

func Login(c *gin.Context){
	var input models.User
	var storedUser models.User
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid data input"})
		return 
	}

	err := authenticateUser(input , storedUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Email or Password is incorrect"})
	}

	token,err := utils.GenerateToken(storedUser.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error generating token"})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"Login Successful",
		"token":token,
	})

}

func insertIntoDb(user models.User, hashedPassword string) error {
	query := `INSERT INTO users(first_name,last_name,email,password,role,gender,age,trainer_code,created_at) 
	VALUES (?,?,?,?,?,?,?,?,NOW())`
	_, err := db.DB.Exec(query,
		user.FirstName,
		user.LastName,
		user.Email,
		string(hashedPassword),
		user.Role,
		user.Gender,
		user.Age,
		user.TrainerCode,
	)
	if err != nil {
		return err
	}
	return nil

}

func authenticateUser(input models.User, storedUser models.User) error {
	query := `SELECT user_id,password FROM users WHERE email = ?`
	err := db.DB.QueryRow(query,input.Email).Scan(&storedUser.UserId,&storedUser.Password)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password),[]byte(input.Password))
	if err != nil {
		return err 
	}
	return nil
}
