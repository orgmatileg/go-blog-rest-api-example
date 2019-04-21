package subscribe

import "github.com/orgmatileg/go-blog-rest-api-example/module/subscribe/model"

// Repository interface
type Repository interface {
	Save(*model.Subscribe) error
	FindByID(id string) (*model.Subscribe, error)
	FindAll(limit, offset, order string) (model.SubscribeList, error)
	Count() (int64, error)
}
