package main

import (
	"log"
	"os"

	"github.com/JatinSihag/Trakr/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err:= godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}
	db.ConnectToDB()

	r:=gin.Default()
	port:=os.Getenv("PORT")
	if port == "" {
		log.Fatal("Error getting port from env file")
	}
	r.Run(":"+port)
}