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
	return u.tagsRepo.Save(mt)
}

func (u *tagsUsecase) FindAll(limit, offset string) (mtl model.Tags, count int64, err error) {
	mtl, err = u.tagsRepo.FindAll(limit, offset)
	if err != nil {
		return nil, -1, nil
	}

	count, err = u.tagsRepo.Count()
	return
}

func (u *tagsUsecase) Delete(tagName string) (err error) {
	return u.tagsRepo.Delete(tagName)
}

func (u *tagsUsecase) Count() (count int64, err error) {
	return u.tagsRepo.Count()
}
