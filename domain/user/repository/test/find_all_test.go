package repository_test

import (
	mocket "github.com/Selvatico/go-mocket"
	"gorm.io/gorm"
)

func (s *RepositorySuite) TestFindAll() {
	reply := []map[string]interface{}{{"ID": 1}}
	mocket.Catcher.NewMock().
		WithQuery(`SELECT * FROM "users"`).
		WithReply(reply)

	users, err := s.repository.FindAll()

	s.Assert().NoError(err)
	s.Assert().NotNil(users)
	s.Assert().EqualValues(len(reply), len(users))
}

func (s *RepositorySuite) TestFindAllWithError() {
	mocket.Catcher.NewMock().
		WithQuery(`SELECT * FROM "users"`).
		WithError(gorm.ErrNotImplemented)

	_, err := s.repository.FindAll()

	s.Assert().Equal(gorm.ErrNotImplemented, err)
}
