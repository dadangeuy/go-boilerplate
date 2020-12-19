package component

import (
	"go-boilerplate/domain/user"
	"go-boilerplate/domain/user/delivery"
	"go-boilerplate/domain/user/repository"
	"go-boilerplate/domain/user/usecase"
	"gorm.io/gorm"
)

func NewUserDelivery(db *gorm.DB) user.Delivery {
	r := &repository.Repository{DB: db}
	u := &usecase.Usecase{Repository: r}
	d := &delivery.Delivery{Usecase: u}

	return d
}
