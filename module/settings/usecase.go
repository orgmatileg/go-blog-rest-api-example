package settings

import "github.com/orgmatileg/go-blog-rest-api-example/module/settings/model"

type Usecase interface {
	FindByID(id string) (*model.Setting, error)
	FindAll(limit, offset string) (msl model.Settings, count int64, err error)
	Update(id string, settingModel *model.Setting) (*string, error)
}
