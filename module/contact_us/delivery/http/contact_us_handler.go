package http

import (
	"encoding/json"
	"hacktiv8/final/helper"
	"hacktiv8/final/module/contact_us"
	"hacktiv8/final/module/contact_us/model"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpContactUsHandler struct {
	CUsecase contact_us.Usecase
}

func NewContactUsHttpHandler(r *mux.Router, cu contact_us.Usecase) {

	handler := HttpContactUsHandler{
		CUsecase: cu,
	}

	r.HandleFunc("/contact-us", handler.ContactUsSaveHttpHandler).Methods("POST")
	r.HandleFunc("/contact-us", handler.ContactUsFindAllHttpHandler).Methods("GET")
	r.HandleFunc("/contact-us/{id}", handler.ContactUsFindByIDHttpHandler).Methods("GET")
	r.HandleFunc("/contact-us/{id}", handler.ContactUsDeleteHttpHandler).Methods("DELETE")
}

// ContactUsSaveHttpHandler handler
func (u *HttpContactUsHandler) ContactUsSaveHttpHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	mc := model.NewContactUs()

	err := decoder.Decode(mc)

	res := helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	err = u.CUsecase.Save(mc)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mc

}

// ContactUsFindByIDHttpHandler handler
func (u *HttpContactUsHandler) ContactUsFindByIDHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	mc, err := u.CUsecase.FindByID(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mc

}

// ContactUsFindAllHttpHandler handler
func (u *HttpContactUsHandler) ContactUsFindAllHttpHandler(w http.ResponseWriter, r *http.Request) {

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

	mcl, err := u.CUsecase.FindAll(limit, offset, order)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mcl

}

// ContactUsDeleteHttpHandler handler
func (u *HttpContactUsHandler) ContactUsDeleteHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idP := vars["id"]

	res := helper.Response{}

	err := u.CUsecase.Delete(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = "OK"

}
