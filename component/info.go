package component

import (
	"go-boilerplate/domain/info/delivery"
	"go-boilerplate/domain/info/repository"
	"go-boilerplate/domain/info/usecase"
	"gorm.io/gorm"
)

func NewInfoDelivery(db *gorm.DB) *delivery.Delivery {
	infoRepository := &repository.Repository{DB: db}
	infoUsecase := &usecase.Usecase{Repository: infoRepository}
	infoDelivery := &delivery.Delivery{Usecase: infoUsecase}

	return infoDelivery
}
