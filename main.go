package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/orgmatileg/go-blog-rest-api-example/config"
	"github.com/orgmatileg/go-blog-rest-api-example/router"

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

	originCORS := handlers.AllowedOrigins([]string{"*"})
	headersCORS := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methodCORS := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(originCORS, headersCORS, methodCORS)(router)))

}
