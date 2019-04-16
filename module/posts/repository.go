package posts

import "github.com/orgmatileg/go-blog-rest-api-example/module/posts/model"

// Repository interface
type Repository interface {
	Save(*model.Post) error
	FindByID(id string) (*model.Post, error)
	// FindAll(limit, offset, order string) (model.ExampleList, error)
	// Update(id string, modelUser *model.Example) (*string, error)
	// Delete(id string) error
	// IsExistsByID(id string) (bool, error)
}
