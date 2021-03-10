package user

import (
	"go-boilerplate/domain/user/delivery"
	"go-boilerplate/domain/user/repository"
	"go-boilerplate/domain/user/usecase"
	"gorm.io/gorm"
)

func NewUserDelivery(usecase Usecase) Delivery {
	return &delivery.Delivery{Usecase: usecase}
}

func NewUserUsecase(repository Repository) Usecase {
	return &usecase.Usecase{Repository: repository}
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repository.Repository{DB: db}
}
