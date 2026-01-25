package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectToDB(){
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("Error:DB_DSN not found in env file")
	}
	var err error
	DB, err = sql.Open("mysql",dsn)
	if err != nil {
		log.Fatal("Error opening db : ",err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Error Connecting to db : ",err)
	}
	fmt.Println("Connected to TrackrDB successfully")
}
func main(){

}