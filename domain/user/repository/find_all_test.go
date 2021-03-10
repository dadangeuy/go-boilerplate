package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
	"regexp"
)

var findAllQuery = regexp.QuoteMeta(`SELECT * FROM "users"`)

func (s *RepositorySuite) TestFindAll() {
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	s.mSQL.
		ExpectQuery(findAllQuery).
		WillReturnRows(rows)

	users, err := s.repository.FindAll()

	s.Assert().NoError(err)
	s.Assert().Len(users, 1)
}

func (s *RepositorySuite) TestFindAllWithError() {
	s.mSQL.
		ExpectQuery(findAllQuery).
		WillReturnError(gorm.ErrNotImplemented)

	_, err := s.repository.FindAll()

	s.Assert().Equal(gorm.ErrNotImplemented, err)
}
