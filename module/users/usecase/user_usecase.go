package usecase

import (
	"github.com/orgmatileg/go-blog-rest-api-example/helper"
	"github.com/orgmatileg/go-blog-rest-api-example/module/users"
	"github.com/orgmatileg/go-blog-rest-api-example/module/users/model"

	"golang.org/x/crypto/bcrypt"
)

type usersUsercase struct {
	usersRepo users.Repository
}

func NewUsersUsecase(u users.Repository) users.Usecase {
	return &usersUsercase{
		usersRepo: u,
	}
}

func (u *usersUsercase) Save(mu *model.User) (err error) {

	// Handle Photo Profile
	defaultPhotoProfile := "https://i.ibb.co/whHn3mf/default-photo-profile.png"

	if mu.PhotoProfile != "" {
		imgBB := helper.NewImgBBConn()
		imgURL, err := imgBB.Upload(mu.PhotoProfile)

		if err != nil {
			return err
		}

		mu.PhotoProfile = imgURL
	} else {
		mu.PhotoProfile = defaultPhotoProfile
	}

	// Handle Password - hashmac
	hash, err := bcrypt.GenerateFromPassword([]byte(mu.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	mu.Password = string(hash)

	err = u.usersRepo.Save(mu)

	return err
}

func (u *usersUsercase) FindByID(idUser string) (mu *model.User, err error) {

	mu, err = u.usersRepo.FindByID(idUser)

	if err != nil {
		return nil, err
	}

	// Clean password, so password not exposed to public
	mu.Password = ""

	return mu, nil
}

func (u *usersUsercase) FindAll(limit, offset, order string) (mul model.Users, count int64, err error) {

	mul, err = u.usersRepo.FindAll(limit, offset, order)

	if err != nil {
		return nil, -1, err
	}

	count, err = u.usersRepo.Count()

	if err != nil {
		return nil, -1, err
	}

	// Clean password, so password not exposed to public
	for _, v := range mul {
		v.Password = ""
	}

	return
}

func (u *usersUsercase) Update(idUser string, mu *model.User) (rowAffected *string, err error) {

	v, err := u.usersRepo.FindByID(idUser)

	if err != nil {
		return nil, err
	}

	// Handle not change photo profile (empty)
	if mu.PhotoProfile == "" {
		mu.PhotoProfile = v.PhotoProfile
	}

	// Handle photo profile is new
	if mu.PhotoProfile != v.PhotoProfile {
		imgBB := helper.NewImgBBConn()
		imgURL, err := imgBB.Upload(mu.PhotoProfile)

		if err != nil {
			return nil, err
		}

		mu.PhotoProfile = imgURL
	}

	// Handle Password
	if mu.Password != "" {
		// create hashmac for the password
		hash, err := bcrypt.GenerateFromPassword([]byte(mu.Password), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}
		mu.Password = string(hash)
	} else {
		// if password is empty, use password user from db
		mu.Password = v.Password
	}

	rowAffected, err = u.usersRepo.Update(idUser, mu)

	if err != nil {
		return nil, err
	}

	return rowAffected, err
}

func (u *usersUsercase) Delete(idUser string) (err error) {

	err = u.usersRepo.Delete(idUser)

	return err
}

func (u *usersUsercase) IsExistsByID(idUser string) (isExist bool, err error) {
	return u.usersRepo.IsExistsByID(idUser)
}

func (u *usersUsercase) Count() (count int64, err error) {
	return u.usersRepo.Count()
}
