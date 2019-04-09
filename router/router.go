package router

import (
	"fmt"
	"hacktiv8/final/config"
	m "hacktiv8/final/middleware"
	h "hacktiv8/final/module/users/delivery/http"
	_usersRepo "hacktiv8/final/module/users/repository"
	_usersUcase "hacktiv8/final/module/users/usecase"
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
