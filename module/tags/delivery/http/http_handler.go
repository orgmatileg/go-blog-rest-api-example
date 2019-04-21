package http

import (
	"encoding/json"
	"net/http"

	"github.com/orgmatileg/go-blog-rest-api-example/helper"
	"github.com/orgmatileg/go-blog-rest-api-example/module/tags"
	"github.com/orgmatileg/go-blog-rest-api-example/module/tags/model"

	"github.com/gorilla/mux"
)

// TagsHandler struct
type TagsHandler struct {
	TUsecase tags.Usecase
}

// NewTagsHTTPHandler func
func NewTagsHTTPHandler(r *mux.Router, tu tags.Usecase) {

	handler := TagsHandler{
		TUsecase: tu,
	}

	r.HandleFunc("/posts/{id}/tags", handler.TagsSaveHTTPHandler).Methods("POST")
	r.HandleFunc("/tags", handler.TagsFindAllHTTPHandler).Methods("GET")
	r.HandleFunc("/posts/{id}/tags", handler.TagsDeleteHTTPHandler).Methods("DELETE")
}

// TagsSaveHTTPHandler handler
func (u *TagsHandler) TagsSaveHTTPHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var mt model.Tag

	err := decoder.Decode(&mt)

	res := helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	err = u.TUsecase.Save(&mt)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mt

}

// TagsFindAllHTTPHandler handler
func (u *TagsHandler) TagsFindAllHTTPHandler(w http.ResponseWriter, r *http.Request) {

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

	res.Body.Payload, res.Body.Count, res.Err = u.TUsecase.FindAll(limit, offset)

	res.ServeJSON(w, r)

}

// TagsDeleteHTTPHandler handler
func (u *TagsHandler) TagsDeleteHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idP := vars["id"]

	res := helper.Response{}

	err := u.TUsecase.Delete(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = "OK"

}
