package users

import "hacktiv8/final/module/users/model"

// Repository interface
type Repository interface {
	Save(*model.User) error
	FindByID(string) (*model.User, error)
	FindAll(string, string, string) (model.Users, error)
	Update(string, *model.User) (*string, error)
	Delete(string) error
}
