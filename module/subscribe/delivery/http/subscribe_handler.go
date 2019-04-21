package http

import (
	"encoding/json"
	"net/http"

	"github.com/orgmatileg/go-blog-rest-api-example/helper"
	"github.com/orgmatileg/go-blog-rest-api-example/module/subscribe"
	"github.com/orgmatileg/go-blog-rest-api-example/module/subscribe/model"

	"github.com/gorilla/mux"
)

type SubscribeHandler struct {
	SUsecase subscribe.Usecase
}

func NewSubscribeHTTPHandler(r *mux.Router, su subscribe.Usecase) {

	handler := SubscribeHandler{
		SUsecase: su,
	}

	r.HandleFunc("/subscribe", handler.SubscribeSaveHTTPHandler).Methods("POST")
	r.HandleFunc("/subscribe", handler.SubscribeFindAllHTTPHandler).Methods("GET")
	r.HandleFunc("/subscribe/{id}", handler.SubscribeFindByIDHTTPHandler).Methods("GET")
}

// SubscribeSaveHTTPHandler handler
func (u *SubscribeHandler) SubscribeSaveHTTPHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	mu := model.NewSubscribe()

	err := decoder.Decode(mu)

	res := helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	err = u.SUsecase.Save(mu)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mu

}

// SubscribeFindByIDHTTPHandler handler
func (u *SubscribeHandler) SubscribeFindByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	mu, err := u.SUsecase.FindByID(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mu

}

// SubscribeFindAllHTTPHandler handler
func (u *SubscribeHandler) SubscribeFindAllHTTPHandler(w http.ResponseWriter, r *http.Request) {

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

	subscribeList, err := u.SUsecase.FindAll(limit, offset, order)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = subscribeList

}
