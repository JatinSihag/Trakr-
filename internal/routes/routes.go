package routes

import (
	"net/http"

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
		protected.GET("/dashboard", func(c *gin.Context) {
			userId, _ := c.Get("userId")
			c.JSON(http.StatusOK, gin.H{"message": "you are authorizied,", "user_Id": userId})
		})
		protected.POST("/profile", handlers.UpdateProfile)
		protected.GET("/foods/search", handlers.SearchFood)
		protected.POST("/foods/log", handlers.LogFood)
		protected.GET("/foods/log", handlers.GetDailyLog)
	}
}
