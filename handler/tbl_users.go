package handler

import (
	"hacktiv8/final/config"
	"hacktiv8/final/helper"
	"hacktiv8/final/repository"
	"net/http"

	"github.com/gorilla/mux"
)

// UserFindByID handler
func UserFindByID(w http.ResponseWriter, r *http.Request) {

	db := config.GetMySQLDB()

	defer db.Close()

	vars := mux.Vars(r)
	res := helper.Response{}

	userRepo := repository.NewUserRepositoryMysql(db)

	user, err := userRepo.FindByID(vars["id"])

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
	}

	res.Body.Payload = user

	return
}

// UserFindAll handler
func UserFindAll(w http.ResponseWriter, r *http.Request) {

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

	return

}
