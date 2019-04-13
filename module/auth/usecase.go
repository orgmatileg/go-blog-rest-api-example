package auth

import (
	modelAuth "hacktiv8/final/module/auth/model"
	"hacktiv8/final/module/users/model"
)

type Usecase interface {
	Login(*model.User) (*modelAuth.Auth, error)
}
