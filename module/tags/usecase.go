package tags

import "github.com/orgmatileg/go-blog-rest-api-example/module/tags/model"

type Usecase interface {
	Save(*model.Tag) error
	FindAll(limit, offset string) (model.Tags, error)
	Delete(tagName string) error
}
