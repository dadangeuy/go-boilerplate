package usecase

import "go-boilerplate/domain/user/repository"

type Usecase struct {
	Repository *repository.Repository
}
