package router

import (
	"encoding/json"
	"fmt"
	"hacktiv8/final/config"
	"hacktiv8/final/helper"
	m "hacktiv8/final/middleware"
	h "hacktiv8/final/module/users/delivery/http"
	_usersRepo "hacktiv8/final/module/users/repository"
	_usersUcase "hacktiv8/final/module/users/usecase"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRouter endpoint
func InitRouter() *mux.Router {

	r := mux.NewRouter()
	// Check API
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!")
	}).Methods("GET")
	// Endpoint for testing app or such a thing
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		t := struct {
			Image string `json:"image"`
		}{}

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println(err)
		}

		err = json.Unmarshal(b, &t)

		if err != nil {
			log.Println(err)
		}

		imgBB := helper.NewImgBBConn()

		imgURL, err := imgBB.Upload(t.Image)

		if err != nil {
			log.Println(err)
		}

		fmt.Println(imgURL)

	}).Methods("POST")

	rv1 := r.PathPrefix("/v1").Subrouter()

	rv1.Use(m.CORS)

	// Get DB Conn
	dbConn := config.GetMySQLDB()

	// Users
	usersRepo := _usersRepo.NewUserRepositoryMysql(dbConn)
	usersUcase := _usersUcase.NewUsersUsecase(usersRepo)
	h.NewUsersHttpHandler(rv1, usersUcase)

	return r
}
