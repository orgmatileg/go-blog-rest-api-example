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

	err = u.exampleRepo.Save(me)

	return err
}

func (u *exampleUsecase) FindByID(id string) (me *model.Example, err error) {

	me, err = u.exampleRepo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return me, nil
}

func (u *exampleUsecase) FindAll(limit, offset, order string) (lme model.ExampleList, err error) {

	lme, err = u.exampleRepo.FindAll(limit, offset, order)

	return lme, err
}

func (u *exampleUsecase) Update(id string, me *model.Example) (rowAffected *string, err error) {

	// v, err := u.exampleRepo.FindByID(id)

	// if err != nil {
	// 	return nil, err
	// }

	rowAffected, err = u.exampleRepo.Update(id, me)

	if err != nil {
		return nil, err
	}

	return rowAffected, err
}

func (u *exampleUsecase) Delete(idUser string) (err error) {

	err = u.exampleRepo.Delete(idUser)

	return err
}

func (u *exampleUsecase) IsExistsByID(idUser string) (isExist bool, err error) {

	isExist, err = u.exampleRepo.IsExistsByID(idUser)

	if err != nil {
		return false, err
	}

	return isExist, nil
}
