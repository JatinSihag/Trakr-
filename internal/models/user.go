package models

import "time"

type User struct {
	UserId      int       `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	Gender      string    `json:"gender"`
	Age         int       `json:"age"`
	TrainerCode string    `json:"trainer_code,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
