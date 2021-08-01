package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go-boilerplate/domain/user/mock"
	"go-boilerplate/domain/user/usecase"
)

type Suite struct {
	suite.Suite
	repository *mock.Repository
	usecase    *usecase.DefaultUsecase
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupTest() {
	s.repository = new(mock.Repository)
	s.usecase = &usecase.DefaultUsecase{Repository: s.repository}
}
