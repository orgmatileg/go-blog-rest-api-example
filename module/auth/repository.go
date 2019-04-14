package auth

import (
	"github.com/orgmatileg/go-blog-rest-api-example/module/users/model"
)

// Repository interface
type Repository interface {
	Login(*model.User) (userModel *model.User, err error)
}
