package usecase

import (
	"go-boilerplate/domain/user/repository"
	"go-boilerplate/model"
)

type Usecase interface {
	List() ([]model.User, error)
}

type DefaultUsecase struct {
	Repository repository.Repository
}
