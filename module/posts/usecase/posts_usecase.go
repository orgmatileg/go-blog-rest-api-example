package usecase

import (
	"github.com/orgmatileg/go-blog-rest-api-example/helper"
	"github.com/orgmatileg/go-blog-rest-api-example/module/posts"
	"github.com/orgmatileg/go-blog-rest-api-example/module/posts/model"
)

type postsUsecase struct {
	postsRepo posts.Repository
}

// NewPostsUsecase NewPostsUsecase
func NewPostsUsecase(pr posts.Repository) posts.Usecase {
	return &postsUsecase{
		postsRepo: pr,
	}
}

func (u *postsUsecase) Save(mp *model.Post) (err error) {

	// Handle Photo Profile
	defaultPhotoProfile := "https://i.ibb.co/pnP96gy/no-image-available.png"

	if mp.PostImage != "" {
		imgBB := helper.NewImgBBConn()
		imgURL, err := imgBB.Upload(mp.PostImage)

		if err != nil {
			return err
		}

		mp.PostImage = imgURL
	} else {
		mp.PostImage = defaultPhotoProfile
	}

	return u.postsRepo.Save(mp)
}

func (u *postsUsecase) FindByID(id string) (me *model.Post, err error) {

	return u.postsRepo.FindByID(id)
}

func (u *postsUsecase) FindAll(limit, offset, order string) (lmp model.Posts, err error) {

	return u.postsRepo.FindAll(limit, offset, order)
}

func (u *postsUsecase) Update(id string, mp *model.Post) (rowAffected *string, err error) {

	rowAffected, err = u.postsRepo.Update(id, mp)

	if err != nil {
		return nil, err
	}

	return rowAffected, err
}

func (u *postsUsecase) Delete(idPost string) (err error) {

	err = u.postsRepo.Delete(idPost)

	return err
}

func (u *postsUsecase) IsExistsByID(idPost string) (isExist bool, err error) {

	isExist, err = u.postsRepo.IsExistsByID(idPost)

	if err != nil {
		return false, err
	}

	return isExist, nil
}
