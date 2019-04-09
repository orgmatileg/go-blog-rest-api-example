package http

import (
	"hacktiv8/final/config"
	"hacktiv8/final/helper"
	"hacktiv8/final/module/users"
	"hacktiv8/final/module/users/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpUsersHandler struct {
	UUsecase users.Usecase
}

func NewUsersHttpHandler(r *mux.Router, uu users.Usecase) {

	handler := HttpUsersHandler{
		UUsecase: uu,
	}

	r.HandleFunc("/users", handler.UserFindAllHttpHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UserFindByIDHttpHandler).Methods("GET")
	// router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
	// router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
}

// UserFindByIDHttpHandler handler
func (u *HttpUsersHandler) UserFindByIDHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	mu, err := u.UUsecase.FindByID(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
	}

	res.Body.Payload = mu

}

// UserFindAllHttpHandler handler
func (u *HttpUsersHandler) UserFindAllHttpHandler(w http.ResponseWriter, r *http.Request) {

	db := config.GetMySQLDB()

	defer db.Close()

	res := helper.Response{}

	userRepo := repository.NewUserRepositoryMysql(db)

	users, err := userRepo.FindAll()

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
	}

	res.Body.Payload = users

}
