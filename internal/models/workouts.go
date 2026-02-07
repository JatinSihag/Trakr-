package models

type WorkoutSet struct {
	SetId      int     `json:"workout_sets_id"`
	ExerciseId int     `json:"exercise_id"`
	SetNumber  int     `json:"set_number"`
	Weight     float64 `json:"weight"`
	Reps       int     `json:"reps"`
	IsPr       bool    `json:"is_pr"`
	Duration   int     `json:"duration_minutes"`
	Calories   int     `json:"set_calories"`
}

type Workouts struct {
	WorkoutId     int    `json:"workout_id"`
	UserId        int    `json:"user_id"`
	WorkoutName   string `json:"workout_name"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	Notes         string `json:"notes"`
	TotalCalories int    `json:"calories_burned"`

	Sets []WorkoutSet `json:"sets"`
}
