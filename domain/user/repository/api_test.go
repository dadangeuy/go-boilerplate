package repository_test

import (
	"github.com/stretchr/testify/suite"
	"go-boilerplate/component"
	"go-boilerplate/domain/user/repository"
	"testing"
)

type RepositorySuite struct {
	suite.Suite
	repository repository.Repository
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (s *RepositorySuite) SetupSuite() {
	db, err := component.NewPostgresDB()
	if err != nil {
		panic(err)
	}

	s.repository.DB = db
}

func (s *RepositorySuite) BeforeTest(_, _ string) {
	s.repository.DB = s.repository.DB.Begin()
}

func (s *RepositorySuite) AfterTest(_, _ string) {
	s.repository.DB = s.repository.DB.Rollback()
}
