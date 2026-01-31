package models

import "time"

type User struct {
	UserId           int       `json:"user_id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	Role             string    `json:"role"`
	Gender           string    `json:"gender"`
	Age              int       `json:"age"`
	Weight           float64   `json:"weight"`
	Height           int       `json:"height"`
	ActivityLevel    string    `json:"activity_level"`
	Goal             string    `json:"goal"`
	TrainerCode      string    `json:"trainer_code,omitempty"`
	VerificationCode string    `json:"-"`
	IsVerified       bool      `json:"is_verified"`
	CodeExpiry       time.Time `json:"-"`
	CreatedAt        time.Time `json:"created_at"`
}

func (u *User) CalculateTDEE() float64 {
 var bmr float64
 heightInFloat := float64(u.Height)
 // Calculate BMR (Mifflin-St Jeor)
 if u.Gender=="male" {
	bmr =(10 *u.Weight)+(6.25*heightInFloat)-(5*float64(u.Age))+5
 }else{
	bmr = (10 *u.Weight)+(6.25*heightInFloat)-(5*float64(u.Age))-161
 }

 multiplier := 1.2
 switch u.ActivityLevel {
 case "light":
	multiplier = 1.375
 case "moderate":
	multiplier = 1.55
 case "active":
	multiplier = 1.725
 case "very_active":
	multiplier = 1.9
 }
 tdee := bmr * multiplier

 var adjustment float64 = 0

 switch u.Goal{
	// for weight lose
 case "lose_0.25":
	adjustment = -275
 case "lose_0.5":
	adjustment = -550
 case "lose_1.0":
	adjustment = -1100
	// for weight gain 
 case "gain_0.25":
	adjustment = 275
 case "gain_0.5":
	adjustment = 550
 case "gain_1.0":
	adjustment = 1100

 case "maintain":
	adjustment = 0
 }
 return tdee+adjustment

}