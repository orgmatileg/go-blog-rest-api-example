package router

import (
	"fmt"

	"github.com/orgmatileg/go-blog-rest-api-example/config"
	m "github.com/orgmatileg/go-blog-rest-api-example/middleware"

	// Auth
	hAuth "github.com/orgmatileg/go-blog-rest-api-example/module/auth/delivery/http"
	_authRepo "github.com/orgmatileg/go-blog-rest-api-example/module/auth/repository"
	_authUcase "github.com/orgmatileg/go-blog-rest-api-example/module/auth/usecase"

	// User
	hUser "github.com/orgmatileg/go-blog-rest-api-example/module/users/delivery/http"
	_usersRepo "github.com/orgmatileg/go-blog-rest-api-example/module/users/repository"
	_usersUcase "github.com/orgmatileg/go-blog-rest-api-example/module/users/usecase"

	// Contact Us
	hContactUs "github.com/orgmatileg/go-blog-rest-api-example/module/contact_us/delivery/http"
	_contactUsRepo "github.com/orgmatileg/go-blog-rest-api-example/module/contact_us/repository"
	_contactUsUcase "github.com/orgmatileg/go-blog-rest-api-example/module/contact_us/usecase"

	// Posts
	"net/http"

	hPosts "github.com/orgmatileg/go-blog-rest-api-example/module/posts/delivery/http"
	_postsRepo "github.com/orgmatileg/go-blog-rest-api-example/module/posts/repository"
	_postsUcase "github.com/orgmatileg/go-blog-rest-api-example/module/posts/usecase"

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

	// Posts
	postsRepo := _postsRepo.NewPostsRepositoryMysql(dbConn)
	postsUcase := _postsUcase.NewPostsUsecase(postsRepo)
	hPosts.NewPostsHTTPHandler(rv1, postsUcase)

	return r
}
