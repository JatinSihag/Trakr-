package routes

import (
	"github.com/JatinSihag/Trakr/internal/handlers"
	"github.com/JatinSihag/Trakr/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	public := r.Group("/v1")
	{
		public.POST("/signup", handlers.SignUp)
		public.POST("/login", handlers.Login)
		public.POST("/forgot-password", handlers.ForgotPassword)
		public.POST("/reset-password", handlers.ResetPassword)
	}
	protected := r.Group("/v1")
	protected.Use(middleware.AuthMiddleWare())
	{
		protected.GET("/dashboard", handlers.GetDashboard)
		protected.POST("/profile", handlers.UpdateProfile)
		protected.GET("/foods/search", handlers.SearchFood)
		protected.POST("/foods/log", handlers.LogFood)
		protected.GET("/foods/log", handlers.GetDailyLog)
		protected.POST("/workouts/log", handlers.LogWorkout)
		protected.GET("/exercises/search", handlers.GetExercise)
	}
}
