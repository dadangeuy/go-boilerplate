package repository_test

import (
	mocket "github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/suite"
	"go-boilerplate/component/mock"
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

	s.repository = repository.Repository{DB: db}
}

func (s *RepositorySuite) SetupTest() {
	mocket.Catcher.Reset()
}
