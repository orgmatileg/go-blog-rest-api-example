package http

import (
	"encoding/json"
	"hacktiv8/final/helper"
	"hacktiv8/final/module/users"
	"hacktiv8/final/module/users/model"
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

	r.HandleFunc("/users", handler.UserSaveHttpHandler).Methods("POST")
	r.HandleFunc("/users", handler.UserFindAllHttpHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UserFindByIDHttpHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UserUpdateHttpHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.UserDeleteHttpHandler).Methods("DELETE")
}

// UserSaveHttpHandler handler
func (u *HttpUsersHandler) UserSaveHttpHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	mu := model.NewUser()

	err := decoder.Decode(mu)

	res := helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	err = u.UUsecase.Save(mu)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mu

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
		return
	}

	res.Body.Payload = mu

}

// UserFindAllHttpHandler handler
func (u *HttpUsersHandler) UserFindAllHttpHandler(w http.ResponseWriter, r *http.Request) {

	queryParam := r.URL.Query()

	// Set default query
	limit := "10"
	offset := "0"
	order := "desc"

	if v := queryParam.Get("limit"); v != "" {
		limit = queryParam.Get("limit")
	}

	if v := queryParam.Get("offset"); v != "" {
		offset = queryParam.Get("offset")
	}

	if v := queryParam.Get("order"); v != "" {
		order = queryParam.Get("order")
	}

	res := helper.Response{}

	users, err := u.UUsecase.FindAll(limit, offset, order)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = users

}

// UserUpdateHttpHandler handler
func (u *HttpUsersHandler) UserUpdateHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	decoder := json.NewDecoder(r.Body)

	mu := model.NewUser()

	err := decoder.Decode(mu)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	err = u.UUsecase.Update(idP, mu)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mu

}

// UserFindByIDHttpHandler handler
func (u *HttpUsersHandler) UserDeleteHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idP := vars["id"]

	res := helper.Response{}

	err := u.UUsecase.Delete(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = "OK"

}
