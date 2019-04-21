package usecase

import (
	"github.com/orgmatileg/go-blog-rest-api-example/module/contact_us"
	"github.com/orgmatileg/go-blog-rest-api-example/module/contact_us/model"
)

type contactUsUsecase struct {
	contactUsRepo contact_us.Repository
}

func NewContactUsUsecase(cr contact_us.Repository) contact_us.Usecase {
	return &contactUsUsecase{
		contactUsRepo: cr,
	}
}

func (u *contactUsUsecase) Save(mc *model.ContactUs) (err error) {
	return u.contactUsRepo.Save(mc)
}

func (u *contactUsUsecase) FindByID(idContact string) (mc *model.ContactUs, err error) {
	return u.contactUsRepo.FindByID(idContact)
}

func (u *contactUsUsecase) FindAll(limit, offset, order string) (mcl model.ContactUsList, count int64, err error) {

	mcl, err = u.contactUsRepo.FindAll(limit, offset, order)

	if err != nil {
		return nil, -1, err
	}

	count, err = u.contactUsRepo.Count()

	if err != nil {
		return nil, -1, err
	}

	return
}

func (u *contactUsUsecase) Delete(idContact string) (err error) {
	return u.contactUsRepo.Delete(idContact)
}

func (u *contactUsUsecase) Count() (int64, error) {
	return u.contactUsRepo.Count()
}
