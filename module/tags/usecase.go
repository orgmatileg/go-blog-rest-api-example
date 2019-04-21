package tags

import "github.com/orgmatileg/go-blog-rest-api-example/module/tags/model"

type Usecase interface {
	Save(*model.Tag) error
	FindAll(limit, offset string) (mtl model.Tags, count int64, err error)
	Delete(tagName string) error
	Count() (int64, error)
}
