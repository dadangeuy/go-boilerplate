package info

import (
	"go-boilerplate/domain/info/delivery"
	"go-boilerplate/domain/info/repository"
	"go-boilerplate/domain/info/usecase"
	"gorm.io/gorm"
)

func NewDelivery(postgres *gorm.DB) *delivery.Delivery {
	infoRepository := &repository.Repository{Postgres: postgres}
	infoUsecase := &usecase.Usecase{Repository: infoRepository}
	infoDelivery := &delivery.Delivery{Usecase: infoUsecase}

	return infoDelivery
}
