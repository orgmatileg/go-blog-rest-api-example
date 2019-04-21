package settings

import "github.com/orgmatileg/go-blog-rest-api-example/module/settings/model"

// Repository interface
type Repository interface {
	FindByID(id string) (*model.Setting, error)
	FindAll(limit, offset string) (mel model.Settings, err error)
	Update(id string, modelUser *model.Setting) (*string, error)
	Count() (int64, error)
}
