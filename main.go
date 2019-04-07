package main

import (
	"hacktiv8/final/model"
	"hacktiv8/final/repository"
	"hacktiv8/final/router"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if env := os.Getenv("GO_ENV"); env != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	router := router.InitRouter()

	log.Fatal(http.ListenAndServe(":8081", router))

}

func findUserAll(repo repository.UserRepository) (model.Users, error) {

	users, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
