package usecase

import (
	"github.com/orgmatileg/go-blog-rest-api-example/module/settings"
	"github.com/orgmatileg/go-blog-rest-api-example/module/settings/model"
)

type settingsUsecase struct {
	settingsRepo settings.Repository
}

func NewSettingsUsecase(sr settings.Repository) settings.Usecase {
	return &settingsUsecase{
		settingsRepo: sr,
	}
}

func (u *settingsUsecase) FindByID(id string) (me *model.Setting, err error) {
	return u.settingsRepo.FindByID(id)
}

func (u *settingsUsecase) FindAll(limit, offset string) (mel model.Settings, count int64, err error) {
	mel, err = u.settingsRepo.FindAll(limit, offset)

	if err != nil {
		return nil, -1, err
	}

	count, err = u.settingsRepo.Count()

	return
}

func (u *settingsUsecase) Update(id string, me *model.Setting) (rowAffected *string, err error) {
	return u.settingsRepo.Update(id, me)
}

func (u *settingsUsecase) Count() (count int64, err error) {
	return u.settingsRepo.Count()
}
