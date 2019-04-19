package usecase

import (
	"github.com/orgmatileg/go-blog-rest-api-example/module/posts"
	"github.com/orgmatileg/go-blog-rest-api-example/module/posts/model"
)

type postsUsecase struct {
	postsRepo posts.Repository
}

func NewPostsUsecase(pr posts.Repository) posts.Usecase {
	return &postsUsecase{
		postsRepo: pr,
	}
}

func (u *postsUsecase) Save(mp *model.Post) (err error) {

	return u.postsRepo.Save(mp)
}

func (u *postsUsecase) FindByID(id string) (me *model.Post, err error) {

	return u.postsRepo.FindByID(id)
}

func (u *postsUsecase) FindAll(limit, offset, order string) (lmp model.Posts, err error) {

	return u.postsRepo.FindAll(limit, offset, order)
}

// func (u *exampleUsecase) Update(id string, me *model.Example) (rowAffected *string, err error) {

// 	// v, err := u.exampleRepo.FindByID(id)

// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	rowAffected, err = u.exampleRepo.Update(id, me)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return rowAffected, err
// }

// func (u *exampleUsecase) Delete(idUser string) (err error) {

// 	err = u.exampleRepo.Delete(idUser)

// 	return err
// }

// func (u *exampleUsecase) IsExistsByID(idUser string) (isExist bool, err error) {

// 	isExist, err = u.exampleRepo.IsExistsByID(idUser)

// 	if err != nil {
// 		return false, err
// 	}

// 	return isExist, nil
// }
