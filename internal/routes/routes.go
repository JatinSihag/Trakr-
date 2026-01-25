package routes

import (
	"github.com/JatinSihag/Trakr/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
 r.POST("/v1/signup",handlers.SignUp)
 r.POST("/v1/login",handlers.Login)
}