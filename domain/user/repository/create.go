package repository

import "go-boilerplate/model"

func (r *DefaultRepository) Create(email string, password string) (model.User, error) {
	user := model.User{
		Email:    email,
		Password: password,
	}
	err := r.DB.Create(&user).Error

	return user, err
}
