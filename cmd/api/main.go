package main

import (
	"log"
	"os"

	"github.com/JatinSihag/Trakr/internal/db"
	"github.com/JatinSihag/Trakr/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}
	db.ConnectToDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Error getting port from env file")
	}
	r.Run(":" + port)
}
