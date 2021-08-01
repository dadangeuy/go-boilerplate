package repository

import (
	"gorm.io/gorm"

	"go-boilerplate/model"
)

type Repository interface {
	Create(email string, password string) (model.User, error)
	FindAll() ([]model.User, error)
}

type DefaultRepository struct {
	DB *gorm.DB
}
