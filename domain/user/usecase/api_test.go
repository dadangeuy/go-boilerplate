package usecase_test

import (
	"github.com/stretchr/testify/suite"
	"go-boilerplate/domain/user/mock"
	"go-boilerplate/domain/user/usecase"
	"testing"
)

type Suite struct {
	suite.Suite
	repository *mock.Repository
	usecase *usecase.Usecase
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupTest() {
	s.repository = new(mock.Repository)
	s.usecase = &usecase.Usecase{Repository: s.repository}
}
