package usecase

import (
	"github.com/orgmatileg/go-blog-rest-api-example/module/tags"
	"github.com/orgmatileg/go-blog-rest-api-example/module/tags/model"
)

type tagsUsecase struct {
	tagsRepo tags.Repository
}

// NewTagsUsecase func
func NewTagsUsecase(tr tags.Repository) tags.Usecase {
	return &tagsUsecase{
		tagsRepo: tr,
	}
}

func (u *tagsUsecase) Save(mt *model.Tag) (err error) {

	err = u.tagsRepo.Save(mt)

	return err
}

func (u *tagsUsecase) FindAll(limit, offset string) (lmt model.Tags, err error) {

	lmt, err = u.tagsRepo.FindAll(limit, offset)

	return lmt, err
}

func (u *tagsUsecase) Delete(tagName string) (err error) {

	err = u.tagsRepo.Delete(tagName)

	return err
}
