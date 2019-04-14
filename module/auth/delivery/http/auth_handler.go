package http

import (
	"encoding/json"
	"github.com/orgmatileg/go-blog-rest-api-example/helper"
	"github.com/orgmatileg/go-blog-rest-api-example/module/auth"
	"github.com/orgmatileg/go-blog-rest-api-example/module/users/model"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpAuthHandler struct {
	AUsecase auth.Usecase
}

func NewAuthHttpHandler(r *mux.Router, au auth.Usecase) {

	handler := HttpAuthHandler{
		AUsecase: au,
	}

	r.HandleFunc("/auth", handler.AuthLoginHttpHandler).Methods("POST")

}

// AuthLoginHttpHandler handler
func (a *HttpAuthHandler) AuthLoginHttpHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var mu model.User

	err := decoder.Decode(&mu)

	res := helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	payload, err := a.AUsecase.Login(&mu)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = payload

}
