package posts

import "github.com/orgmatileg/go-blog-rest-api-example/module/posts/model"

// Repository Posts interface
type Repository interface {
	Save(*model.Post) error
	FindByID(id string) (*model.Post, error)
	FindAll(limit, offset, order, isPublish string) (mpl model.Posts, err error)
	Update(id string, modelPost *model.Post) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
	Count() (int64, error)
}
