package example

import "github.com/orgmatileg/go-blog-rest-api-example/module/example/model"

// Repository interface
type Repository interface {
	Save(*model.Example) error
	FindByID(id string) (*model.Example, error)
	FindAll(limit, offset, order string) (model.ExampleList, error)
	Update(id string, modelUser *model.Example) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
}
