package user

import (
	"gorm.io/gorm"

	"go-boilerplate/domain/user/delivery"
	"go-boilerplate/domain/user/repository"
	"go-boilerplate/domain/user/usecase"
)

func NewUserDelivery(db *gorm.DB) delivery.Delivery {
	r := &repository.DefaultRepository{DB: db}
	u := &usecase.DefaultUsecase{Repository: r}
	d := &delivery.DefaultDelivery{Usecase: u}

	return d
}
