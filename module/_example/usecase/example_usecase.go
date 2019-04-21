package usecase

import (
	"github.com/orgmatileg/go-blog-rest-api-example/module/example"
	"github.com/orgmatileg/go-blog-rest-api-example/module/example/model"
)

type exampleUsecase struct {
	exampleRepo example.Repository
}

func NewExampleUsecase(er example.Repository) example.Usecase {
	return &exampleUsecase{
		exampleRepo: er,
	}
}

func (u *exampleUsecase) Save(me *model.Example) (err error) {
	return u.exampleRepo.Save(me)
}

func (u *exampleUsecase) FindByID(id string) (me *model.Example, err error) {
	return u.exampleRepo.FindByID(id)
}

func (u *exampleUsecase) FindAll(limit, offset, order string) (lme model.ExampleList, err error) {
	return u.exampleRepo.FindAll(limit, offset, order)
}

func (u *exampleUsecase) Update(id string, me *model.Example) (rowAffected *string, err error) {
	return u.exampleRepo.Update(id, me)
}

func (u *exampleUsecase) Delete(idUser string) (err error) {
	return u.exampleRepo.Delete(idUser)
}

func (u *exampleUsecase) IsExistsByID(idUser string) (isExist bool, err error) {
	return u.exampleRepo.IsExistsByID(idUser)
}

func (u *exampleUsecase) Count() (count int64, err error) {
	return u.exampleRepo.Count()
}
