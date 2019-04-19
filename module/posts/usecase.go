package posts

import "github.com/orgmatileg/go-blog-rest-api-example/module/posts/model"

// Usecase Posts Interface
type Usecase interface {
	Save(*model.Post) error
	FindByID(id string) (*model.Post, error)
	FindAll(limit, offset, order, isPublish string) (model.Posts, error)
	Update(id string, modelPost *model.Post) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
}
