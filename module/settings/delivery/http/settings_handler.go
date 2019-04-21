package http

import (
	"encoding/json"
	"net/http"

	"github.com/orgmatileg/go-blog-rest-api-example/helper"
	"github.com/orgmatileg/go-blog-rest-api-example/module/settings"
	"github.com/orgmatileg/go-blog-rest-api-example/module/settings/model"

	"github.com/gorilla/mux"
)

type SettingsHandler struct {
	SUsecase settings.Usecase
}

func NewSettingsHttpHandler(r *mux.Router, su settings.Usecase) {

	handler := SettingsHandler{
		SUsecase: su,
	}

	r.HandleFunc("/settings", handler.SettingsFindAllHTTPHandler).Methods("GET")
	r.HandleFunc("/settings/{id}", handler.SettingsFindByIDHTTPHandler).Methods("GET")
	r.HandleFunc("/settings/{id}", handler.SettingsUpdateHTTPHandler).Methods("PUT")
}

// SettingsFindByIDHTTPHandler handler
func (u *SettingsHandler) SettingsFindByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	res.Body.Payload, res.Err = u.SUsecase.FindByID(idP)

	res.ServeJSON(w, r)

}

// ExampleFindAllHttpHandler handler
func (u *SettingsHandler) SettingsFindAllHTTPHandler(w http.ResponseWriter, r *http.Request) {

	queryParam := r.URL.Query()

	// Set default query
	limit := "10"
	offset := "0"

	if v := queryParam.Get("limit"); v != "" {
		limit = queryParam.Get("limit")
	}

	if v := queryParam.Get("offset"); v != "" {
		offset = queryParam.Get("offset")
	}

	res := helper.Response{}

	res.Body.Payload, res.Body.Count, res.Err = u.SUsecase.FindAll(limit, offset)

	res.ServeJSON(w, r)

}

// ExampleUpdateHttpHandler handler
func (u *SettingsHandler) SettingsUpdateHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	decoder := json.NewDecoder(r.Body)

	var ms model.Setting

	err := decoder.Decode(&ms)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload, res.Err = u.SUsecase.Update(idP, &ms)

}
