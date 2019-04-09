package main

import (
	"fmt"
	"hacktiv8/final/config"
	"hacktiv8/final/model"
	"hacktiv8/final/repository"
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

		fmt.Println("diisini")

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

func findUserAll(repo repository.UserRepository) (model.Users, error) {

	users, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
