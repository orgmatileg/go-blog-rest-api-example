package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/orgmatileg/go-blog-rest-api-example/helper"
	example "github.com/orgmatileg/go-blog-rest-api-example/module/_example"
	"github.com/orgmatileg/go-blog-rest-api-example/module/_example/model"

	"github.com/gorilla/mux"
)

type HttpExampleHandler struct {
	EUsecase example.Usecase
}

func NewExampleHttpHandler(r *mux.Router, eu example.Usecase) {

	handler := HttpExampleHandler{
		EUsecase: eu,
	}

	r.HandleFunc("/example", handler.ExampleSaveHttpHandler).Methods("POST")
	r.HandleFunc("/example", handler.ExampleFindAllHttpHandler).Methods("GET")
	r.HandleFunc("/example/{id}", handler.ExampleFindByIDHttpHandler).Methods("GET")
	r.HandleFunc("/example/{id}", handler.ExampleUpdateHttpHandler).Methods("PUT")
	r.HandleFunc("/example/{id}", handler.ExampleDeleteHttpHandler).Methods("DELETE")
	r.HandleFunc("/example/{id}/exists", handler.ExampleIsExistsByIDHttpHandler).Methods("GET")
}

// ExampleSaveHttpHandler handler
func (u *HttpExampleHandler) ExampleSaveHttpHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	mu := model.NewExample()

	err := decoder.Decode(mu)

	res := helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	err = u.EUsecase.Save(mu)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mu

}

// ExampleFindByIDHttpHandler handler
func (u *HttpExampleHandler) ExampleFindByIDHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	mu, err := u.EUsecase.FindByID(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mu

}

// ExampleFindAllHttpHandler handler
func (u *HttpExampleHandler) ExampleFindAllHttpHandler(w http.ResponseWriter, r *http.Request) {

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

	res.Body.Payload, res.Body.Count, res.Err = u.EUsecase.FindAll(limit, offset, order)

	res.ServeJSON(w, r)

}

// ExampleUpdateHttpHandler handler
func (u *HttpExampleHandler) ExampleUpdateHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	decoder := json.NewDecoder(r.Body)

	var me model.Example

	err := decoder.Decode(&me)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	rowsAffected, err := u.EUsecase.Update(idP, &me)

	if err != nil {
		res.Err = err
		return
	}

	fmt.Println(rowsAffected, err)

	res.Body.Payload = fmt.Sprintf("Total rows affected: %s", *rowsAffected)

}

// ExampleDeleteHttpHandler handler
func (u *HttpExampleHandler) ExampleDeleteHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idP := vars["id"]

	res := helper.Response{}

	err := u.EUsecase.Delete(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = "OK"

}

// ExampleIsExistsByIDHttpHandler handler
func (u *HttpExampleHandler) ExampleIsExistsByIDHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	isExists, err := u.EUsecase.IsExistsByID(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	if isExists {
		res.Body.Payload = "ID Example " + idP + " is Exists!"
	} else {
		res.Body.Payload = "ID Example " + idP + " is Not Exists!"
	}

}
