package usecase

import "go-boilerplate/model"

func (u *Usecase) List() ([]model.User, error) {
	return u.Repository.FindAll()
}
