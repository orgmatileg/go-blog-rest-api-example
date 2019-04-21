package usecase

import (
	"fmt"

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

func (u *postsUsecase) FindAll(limit, offset, order, isPublish string) (mpl model.Posts, count int64, err error) {

	mpl, err = u.postsRepo.FindAll(limit, offset, order, isPublish)
	if err != nil {
		return nil, -1, err
	}
	count, err = u.postsRepo.Count()

	return
}

func (u *postsUsecase) Update(id string, mp *model.Post) (rowAffected *string, err error) {

	fmt.Println(id)

	v, err := u.FindByID(id)

	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	if mp.PostImage != "" {
		imgBB := helper.NewImgBBConn()
		imgURL, err := imgBB.Upload(mp.PostImage)

		if err != nil {
			return nil, err
		}

		v.PostImage = imgURL
	}

	v.PostSubject = mp.PostSubject
	v.PostContent = mp.PostContent
	v.IsPublish = mp.IsPublish
	v.Author = mp.Author

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
	return u.postsRepo.IsExistsByID(idPost)
}

func (u *postsUsecase) Count() (count int64, err error) {
	return u.postsRepo.Count()
}
