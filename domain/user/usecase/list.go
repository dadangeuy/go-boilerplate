package usecase

import "go-boilerplate/model"

func (u *Usecase) FindAll() ([]model.User, error) {
	return u.Repository.FindAll()
}
