package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-boilerplate/domain/user/repository"
)

type RepositorySuite struct {
	suite.Suite
	mSQL       sqlmock.Sqlmock
	repository repository.DefaultRepository
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (s *RepositorySuite) SetupSuite() {

}

func (s *RepositorySuite) SetupTest() {
	mDB, mSQL, err := sqlmock.New()
	s.Assert().NoError(err)

	s.mSQL = mSQL

	config := postgres.Config{Conn: mDB}
	dialect := postgres.New(config)
	db, err := gorm.Open(dialect, nil)
	s.repository = repository.DefaultRepository{DB: db}
}
