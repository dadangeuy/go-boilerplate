package repository

import "go-boilerplate/model"

func (r *Repository) Create(username string, password string) (model.User, error) {
	user := model.User{
		Username: username,
		Password: password,
	}
	err := r.DB.Create(&user).Error

	return user, err
}
