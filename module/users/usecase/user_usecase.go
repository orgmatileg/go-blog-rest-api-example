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

func (u *usersUsercase) FindByID(idUser string) (mu *model.User, err error) {

	mu, err = u.usersRepo.FindByID(idUser)

	if err != nil {
		return nil, err
	}

	return mu, nil
}

func (u *usersUsercase) FindAll() (lmu model.Users, err error) {

	lmu, err = u.usersRepo.FindAll()

	return lmu, err
}
