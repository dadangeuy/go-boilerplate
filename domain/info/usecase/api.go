package usecase

import "go-boilerplate/domain/info/repository"

type Usecase struct {
	Repository *repository.Repository
}
