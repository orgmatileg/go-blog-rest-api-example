package usecase

import (
	"hacktiv8/final/module/users"
	"hacktiv8/final/module/users/model"
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

	err = u.usersRepo.Save(mu)

	return err
}

func (u *usersUsercase) FindByID(idUser string) (mu *model.User, err error) {

	mu, err = u.usersRepo.FindByID(idUser)

	if err != nil {
		return nil, err
	}

	return mu, nil
}

func (u *usersUsercase) FindAll(limit, offset, order string) (lmu model.Users, err error) {

	lmu, err = u.usersRepo.FindAll(limit, offset, order)

	return lmu, err
}

func (u *usersUsercase) Update(idUser string, mu *model.User) (err error) {

	err = u.usersRepo.Update(idUser, mu)

	return err
}

func (u *usersUsercase) Delete(idUser string) (err error) {

	err = u.usersRepo.Delete(idUser)

	return err
}
