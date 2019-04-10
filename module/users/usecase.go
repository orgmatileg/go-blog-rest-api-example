package users

import "hacktiv8/final/module/users/model"

type Usecase interface {
	// Save(*model.User) error
	FindByID(string) (*model.User, error)
	FindAll(string, string, string) (model.Users, error)
	// Update(string, *model.User) error
	// Delete(string) error
}
