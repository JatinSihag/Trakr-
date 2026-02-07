package models

type Exercise struct {
	ExerciseId int `json:"exercise_id"`
	ExerciseName string `json:"exercise_name"`
	BodyPart string `json:"body_part"`
	Category string `json:"category"`
}