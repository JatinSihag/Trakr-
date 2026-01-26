package handlers

import (
	"fmt"
	"net/http"
	"time"

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

func Login(c *gin.Context) {
	var input models.User
	var storedUser models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data input"})
		return
	}

	err := authenticateUser(input, &storedUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or Password is incorrect"})
		return
	}

	token, err := utils.GenerateToken(storedUser.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
		"token":   token,
	})

}

func ForgotPassword(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		return
	}
	var user models.User
	query := "SELECT user_id FROM users WHERE email = ?"
	err := db.DB.QueryRow(query, input.Email).Scan(&user.UserId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "If this email exists, a code has been sent."})
		return
	}
	otp := utils.GenerateOTP()
	expiry := time.Now().Add(15 * time.Minute)

	updateQuery := `UPDATE users SET verification_code = ?,code_expiry= ? WHERE email = ?`
	_, err = db.DB.Exec(updateQuery, otp, expiry, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update db"})
		return
	}
	utils.SendEmail(input.Email, otp)
	c.JSON(http.StatusOK, gin.H{"message": "Reset code sent to email"})
}

func ResetPassword(c *gin.Context) {
	var input struct {
		Email       string `json:"email"`
		OTP         string `json:"otp"`
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var user models.User
	query := `SELECT user_id,code_expiry FROM users WHERE email=? AND verification_code=?`
	err := db.DB.QueryRow(query, input.Email, input.OTP).Scan(&user.UserId, &user.CodeExpiry)
	if err != nil {
		fmt.Println("❌ DATABASE ERROR:", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Code or Email"})
		return
	}
	if time.Now().After(user.CodeExpiry) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Code expired"})
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.NewPassword), 14)

	updateQuery := `UPDATE users SET password = ?,verification_code=NULL,code_expiry=NULL WHERE email=?`
	_, err = db.DB.Exec(updateQuery, hashedPassword, input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed updating DB"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success": "Password reset successfully"})

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

func authenticateUser(input models.User, storedUser *models.User) error {
	query := `SELECT user_id,password FROM users WHERE email = ?`
	err := db.DB.QueryRow(query, input.Email).Scan(&storedUser.UserId, &storedUser.Password)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(input.Password))
	if err != nil {
		return err
	}
	return nil
}
