package subscribe

import "github.com/orgmatileg/go-blog-rest-api-example/module/subscribe/model"

type Usecase interface {
	Save(*model.Subscribe) error
	FindByID(id string) (*model.Subscribe, error)
	FindAll(limit, offset, order string) (msl model.SubscribeList, count int64, err error)
	Count() (int64, error)
}
