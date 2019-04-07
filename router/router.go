package router

import (
	h "hacktiv8/final/handler"
	m "hacktiv8/final/middleware"

	"github.com/gorilla/mux"
)

// InitRouter endpoint
func InitRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/ping", h.Ping).Methods("GET")

	rv1 := r.PathPrefix("/v1").Subrouter()

	rv1.Use(m.CORS)

	rv1.HandleFunc("/users", h.UserFindAll).Methods("GET")
	rv1.HandleFunc("/users/{id}", h.UserFindByID).Methods("GET")
	// router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
	// router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	return r
}
