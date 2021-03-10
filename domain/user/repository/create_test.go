package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
	"regexp"
)

var createQuery = regexp.QuoteMeta(`INSERT INTO "users" ("email","password")`)

func (s *RepositorySuite) TestCreate() {
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	s.mSQL.
		ExpectQuery(createQuery).
		WithArgs("hello", "123456").
		WillReturnRows(rows)

	user, err := s.repository.Create("hello", "123456")

	s.Assert().NoError(err)
	s.Assert().EqualValues(1, user.ID)
	s.Assert().EqualValues("hello", user.Email)
	s.Assert().EqualValues("123456", user.Password)
}

func (s *RepositorySuite) TestCreateWithError() {
	s.mSQL.
		ExpectQuery(createQuery).
		WillReturnError(gorm.ErrNotImplemented)

	_, err := s.repository.Create("hello", "123456")

	s.Assert().Equal(gorm.ErrNotImplemented, err)
}
