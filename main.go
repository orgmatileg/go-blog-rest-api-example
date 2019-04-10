package main

import (
	"hacktiv8/final/config"
	"hacktiv8/final/router"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {

	if env := os.Getenv("GO_ENV"); env != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	config.InitConnectionDB()
}

func main() {

	router := router.InitRouter()

	log.Fatal(http.ListenAndServe(":8081", router))

}
