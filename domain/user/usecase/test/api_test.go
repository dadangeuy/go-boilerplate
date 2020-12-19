package usecase_test

import (
	"github.com/stretchr/testify/suite"
	"go-boilerplate/domain/user/mocks"
	"go-boilerplate/domain/user/usecase"
	"testing"
)

type Suite struct {
	suite.Suite
	repository *mocks.Repository
	usecase *usecase.Usecase
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupTest() {
	s.repository = new(mocks.Repository)
	s.usecase = &usecase.Usecase{Repository: s.repository}
}
