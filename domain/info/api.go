package info

import (
	"go-boilerplate/domain/info/delivery"
	"go-boilerplate/domain/info/repository"
	"go-boilerplate/domain/info/usecase"
)

func NewDelivery() delivery.Delivery {
	return delivery.Delivery{
		Usecase: usecase.Usecase{
			Repository: repository.Repository{},
		},
	}
}
