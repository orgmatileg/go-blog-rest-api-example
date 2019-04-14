package posts

import "github.com/orgmatileg/go-blog-rest-api-example/module/posts/model"

type Usecase interface {
	Save(*model.Post) error
	// FindByID(id string) (*model.Example, error)
	// FindAll(limit, offset, order string) (model.ExampleList, error)
	// Update(id string, exampleModel *model.Example) (*string, error)
	// Delete(id string) error
	// IsExistsByID(id string) (bool, error)
}
