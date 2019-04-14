package contact_us

import "github.com/orgmatileg/go-blog-rest-api-example/module/contact_us/model"

type Usecase interface {
	Save(*model.ContactUs) error
	FindByID(id string) (*model.ContactUs, error)
	FindAll(limit, offset, order string) (model.ContactUsList, error)
	Delete(id string) error
}
