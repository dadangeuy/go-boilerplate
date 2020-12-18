package repository

import "go-boilerplate/model"

func (r *Repository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error

	return users, err
}
