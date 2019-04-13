package auth

import (
	"hacktiv8/final/module/users/model"
)

// Repository interface
type Repository interface {
	Login(*model.User) (userModel *model.User, err error)
}
