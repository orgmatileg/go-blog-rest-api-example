package users

import "hacktiv8/final/module/users/model"

// Repository interface
type Repository interface {
	Save(*model.User) error
	FindByID(id string) (*model.User, error)
	FindAll(limit, offset, order string) (model.Users, error)
	Update(id string, modelUser *model.User) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
}
