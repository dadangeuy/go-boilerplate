package usecase

import "go-boilerplate/model"

func (u *DefaultUsecase) List() ([]model.User, error) {
	return u.Repository.FindAll()
}
