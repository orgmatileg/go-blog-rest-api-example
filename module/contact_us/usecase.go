package contact_us

import "hacktiv8/final/module/contact_us/model"

type Usecase interface {
	Save(*model.ContactUs) error
	FindByID(id string) (*model.ContactUs, error)
	FindAll(limit, offset, order string) (model.ContactUsList, error)
	Delete(id string) error
}
