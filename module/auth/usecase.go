package auth

import (
	modelAuth "github.com/orgmatileg/go-blog-rest-api-example/module/auth/model"
	"github.com/orgmatileg/go-blog-rest-api-example/module/users/model"
)

type Usecase interface {
	Login(*model.User) (*modelAuth.Auth, error)
}
