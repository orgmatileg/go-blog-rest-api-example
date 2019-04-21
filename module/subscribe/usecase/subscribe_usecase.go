package usecase

import (
	"github.com/orgmatileg/go-blog-rest-api-example/module/subscribe"
	"github.com/orgmatileg/go-blog-rest-api-example/module/subscribe/model"
)

type subscribeUsecase struct {
	subscribeRepo subscribe.Repository
}

func NewSubscribeUsecase(sr subscribe.Repository) subscribe.Usecase {
	return &subscribeUsecase{
		subscribeRepo: sr,
	}
}

func (u *subscribeUsecase) Save(me *model.Subscribe) (err error) {
	return u.subscribeRepo.Save(me)
}

func (u *subscribeUsecase) FindByID(id string) (ms *model.Subscribe, err error) {
	return u.subscribeRepo.FindByID(id)
}

func (u *subscribeUsecase) FindAll(limit, offset, order string) (msl model.SubscribeList, err error) {
	return u.subscribeRepo.FindAll(limit, offset, order)
}
