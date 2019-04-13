package router

import (
	"fmt"
	"hacktiv8/final/config"
	m "hacktiv8/final/middleware"
	hAuth "hacktiv8/final/module/auth/delivery/http"
	_authRepo "hacktiv8/final/module/auth/repository"
	_authUcase "hacktiv8/final/module/auth/usecase"
	hUser "hacktiv8/final/module/users/delivery/http"
	_usersRepo "hacktiv8/final/module/users/repository"
	_usersUcase "hacktiv8/final/module/users/usecase"

	hContactUs "hacktiv8/final/module/contact_us/delivery/http"
	_contactUsRepo "hacktiv8/final/module/contact_us/repository"
	_contactUsUcase "hacktiv8/final/module/contact_us/usecase"
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

		fmt.Fprintf(w, "Test!")
	}).Methods("POST")

	// Init versioning API
	rv1 := r.PathPrefix("/v1").Subrouter()

	// Middleware
	rv1.Use(m.CheckAuth)
	rv1.Use(m.CORS)

	// Get DB Conn
	dbConn := config.GetMySQLDB()

	// Auth
	authRepo := _authRepo.NewAuthRepositoryMysql(dbConn)
	authUcase := _authUcase.NewAuthUsecase(authRepo)
	hAuth.NewAuthHttpHandler(rv1, authUcase)

	// Users
	usersRepo := _usersRepo.NewUserRepositoryMysql(dbConn)
	usersUcase := _usersUcase.NewUsersUsecase(usersRepo)
	hUser.NewUsersHttpHandler(rv1, usersUcase)

	// Contact Us
	contactUsRepo := _contactUsRepo.NewContactUsRepositoryMysql(dbConn)
	contactUsUcase := _contactUsUcase.NewContactUsUsecase(contactUsRepo)
	hContactUs.NewContactUsHttpHandler(rv1, contactUsUcase)

	return r
}
