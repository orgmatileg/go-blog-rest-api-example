package http

import (
	"encoding/json"
	"net/http"

	"github.com/orgmatileg/go-blog-rest-api-example/helper"
	"github.com/orgmatileg/go-blog-rest-api-example/module/posts"
	"github.com/orgmatileg/go-blog-rest-api-example/module/posts/model"

	"github.com/gorilla/mux"
)

type HttpPostsHandler struct {
	PUsecase posts.Usecase
}

func NewPostsHttpHandler(r *mux.Router, pu posts.Usecase) {

	handler := HttpPostsHandler{
		PUsecase: pu,
	}

	r.HandleFunc("/posts", handler.PostsSaveHttpHandler).Methods("POST")
	// r.HandleFunc("/example", handler.ExampleFindAllHttpHandler).Methods("GET")
	// r.HandleFunc("/example/{id}", handler.ExampleFindByIDHttpHandler).Methods("GET")
	// r.HandleFunc("/example/{id}", handler.ExampleUpdateHttpHandler).Methods("PUT")
	// r.HandleFunc("/example/{id}", handler.ExampleDeleteHttpHandler).Methods("DELETE")
	// r.HandleFunc("/example/{id}/exists", handler.ExampleIsExistsByIDHttpHandler).Methods("GET")
}

// PostsSaveHttpHandler handler
func (u *HttpPostsHandler) PostsSaveHttpHandler(w http.ResponseWriter, r *http.Request) {

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

// // ExampleFindByIDHttpHandler handler
// func (u *HttpExampleHandler) ExampleFindByIDHttpHandler(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	res := helper.Response{}

// 	idP := vars["id"]

// 	mu, err := u.EUsecase.FindByID(idP)

// 	defer res.ServeJSON(w, r)

// 	if err != nil {
// 		res.Err = err
// 		return
// 	}

// 	res.Body.Payload = mu

// }

// // ExampleFindAllHttpHandler handler
// func (u *HttpExampleHandler) ExampleFindAllHttpHandler(w http.ResponseWriter, r *http.Request) {

// 	queryParam := r.URL.Query()

// 	// Set default query
// 	limit := "10"
// 	offset := "0"
// 	order := "desc"

// 	if v := queryParam.Get("limit"); v != "" {
// 		limit = queryParam.Get("limit")
// 	}

// 	if v := queryParam.Get("offset"); v != "" {
// 		offset = queryParam.Get("offset")
// 	}

// 	if v := queryParam.Get("order"); v != "" {
// 		order = queryParam.Get("order")
// 	}

// 	res := helper.Response{}

// 	Examples, err := u.EUsecase.FindAll(limit, offset, order)

// 	defer res.ServeJSON(w, r)

// 	if err != nil {
// 		res.Err = err
// 		return
// 	}

// 	res.Body.Payload = Examples

// }

// // ExampleUpdateHttpHandler handler
// func (u *HttpExampleHandler) ExampleUpdateHttpHandler(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	res := helper.Response{}

// 	idP := vars["id"]

// 	decoder := json.NewDecoder(r.Body)

// 	var me model.Example

// 	err := decoder.Decode(&me)

// 	defer res.ServeJSON(w, r)

// 	if err != nil {
// 		res.Err = err
// 		return
// 	}

// 	rowsAffected, err := u.EUsecase.Update(idP, &me)

// 	if err != nil {
// 		res.Err = err
// 		return
// 	}

// 	fmt.Println(rowsAffected, err)

// 	res.Body.Payload = fmt.Sprintf("Total rows affected: %s", *rowsAffected)

// }

// // ExampleDeleteHttpHandler handler
// func (u *HttpExampleHandler) ExampleDeleteHttpHandler(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)

// 	idP := vars["id"]

// 	res := helper.Response{}

// 	err := u.EUsecase.Delete(idP)

// 	defer res.ServeJSON(w, r)

// 	if err != nil {
// 		res.Err = err
// 		return
// 	}

// 	res.Body.Payload = "OK"

// }

// // ExampleIsExistsByIDHttpHandler handler
// func (u *HttpExampleHandler) ExampleIsExistsByIDHttpHandler(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	res := helper.Response{}

// 	idP := vars["id"]

// 	isExists, err := u.EUsecase.IsExistsByID(idP)

// 	defer res.ServeJSON(w, r)

// 	if err != nil {
// 		res.Err = err
// 		return
// 	}

// 	if isExists {
// 		res.Body.Payload = "ID Example " + idP + " is Exists!"
// 	} else {
// 		res.Body.Payload = "ID Example " + idP + " is Not Exists!"
// 	}

// }
