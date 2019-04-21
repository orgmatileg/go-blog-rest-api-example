package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/orgmatileg/go-blog-rest-api-example/helper"
	"github.com/orgmatileg/go-blog-rest-api-example/module/posts"
	"github.com/orgmatileg/go-blog-rest-api-example/module/posts/model"

	"github.com/gorilla/mux"
)

// PostsHandler struct
type PostsHandler struct {
	PUsecase posts.Usecase
}

// NewPostsHTTPHandler func
func NewPostsHTTPHandler(r *mux.Router, pu posts.Usecase) {

	handler := PostsHandler{
		PUsecase: pu,
	}

	r.HandleFunc("/posts", handler.PostsSaveHTTPHandler).Methods("POST")
	r.HandleFunc("/posts", handler.PostsFindAllHTTPHandler).Methods("GET")
	r.HandleFunc("/posts/{id}", handler.PostsFindByIDHTTPHandler).Methods("GET")
	r.HandleFunc("/posts/{id}", handler.PostsUpdateHTTPHandler).Methods("PUT")
	r.HandleFunc("/posts/{id}", handler.PostsDeleteHTTPHandler).Methods("DELETE")
	r.HandleFunc("/posts/{id}/exists", handler.PostsIsExistsByIDHTTPHandler).Methods("GET")
}

// PostsSaveHTTPHandler handler
func (u *PostsHandler) PostsSaveHTTPHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	mp := model.NewPost()

	err := decoder.Decode(mp)

	res := helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	err = u.PUsecase.Save(mp)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mp

}

// PostsFindByIDHTTPHandler handler
func (u *PostsHandler) PostsFindByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	mu, err := u.PUsecase.FindByID(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mu

}

// PostsFindAllHTTPHandler handler
func (u *PostsHandler) PostsFindAllHTTPHandler(w http.ResponseWriter, r *http.Request) {

	queryParam := r.URL.Query()

	// Set default query
	limit := "10"
	offset := "0"
	order := "desc"
	isPublish := "1"

	if v := queryParam.Get("limit"); v != "" {
		limit = queryParam.Get("limit")
	}

	if v := queryParam.Get("offset"); v != "" {
		offset = queryParam.Get("offset")
	}

	if v := queryParam.Get("order"); v != "" {
		order = queryParam.Get("order")
	}

	if v := queryParam.Get("isPublish"); v != "" {
		isPublish = queryParam.Get("isPublish")
	}

	res := helper.Response{}

	res.Body.Payload, res.Body.Count, res.Err = u.PUsecase.FindAll(limit, offset, order, isPublish)

	defer res.ServeJSON(w, r)

}

// PostsUpdateHTTPHandler handler
func (u *PostsHandler) PostsUpdateHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	decoder := json.NewDecoder(r.Body)

	var mp model.Post

	err := decoder.Decode(&mp)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	rowsAffected, err := u.PUsecase.Update(idP, &mp)

	if err != nil {
		res.Err = err
		return
	}

	fmt.Println(rowsAffected, err)

	res.Body.Payload = fmt.Sprintf("Total rows affected: %s", *rowsAffected)

}

// PostsDeleteHTTPHandler handler
func (u *PostsHandler) PostsDeleteHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idP := vars["id"]

	res := helper.Response{}

	err := u.PUsecase.Delete(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = "OK"

}

// PostsIsExistsByIDHTTPHandler handler
func (u *PostsHandler) PostsIsExistsByIDHTTPHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	isExists, err := u.PUsecase.IsExistsByID(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	if isExists {
		res.Body.Payload = "ID Posts " + idP + " is Exists!"
	} else {
		res.Body.Payload = "ID Posts " + idP + " is Not Exists!"
	}

}
