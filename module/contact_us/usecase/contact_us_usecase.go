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

	err = u.contactUsRepo.Save(mc)

	return err
}

func (u *contactUsUsecase) FindByID(idContact string) (mc *model.ContactUs, err error) {

	mc, err = u.contactUsRepo.FindByID(idContact)

	if err != nil {
		return nil, err
	}

	return mc, nil
}

func (u *contactUsUsecase) FindAll(limit, offset, order string) (mcl model.ContactUsList, err error) {

	mcl, err = u.contactUsRepo.FindAll(limit, offset, order)

	return mcl, err
}

func (u *contactUsUsecase) Delete(idContact string) (err error) {

	err = u.contactUsRepo.Delete(idContact)

	return err
}
