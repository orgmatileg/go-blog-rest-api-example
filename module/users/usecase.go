package users

import "github.com/orgmatileg/go-blog-rest-api-example/module/users/model"

type Usecase interface {
	Save(*model.User) error
	FindByID(id string) (*model.User, error)
	FindAll(limit, offset, order string) (mul model.Users, count int64, err error)
	Update(id string, modelUser *model.User) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
	Count() (int64, error)
}
