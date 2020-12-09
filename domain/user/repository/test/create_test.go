package repository_test

import (
	mocket "github.com/Selvatico/go-mocket"
	"gorm.io/gorm"
)

func (s *RepositorySuite) TestCreate() {
	reply := []map[string]interface{}{{"ID": 1}}
	mocket.Catcher.NewMock().
		WithQuery(`INSERT INTO "users"`).
		WithReply(reply)

	user, err := s.repository.Create("hello", "123456")

	s.Assert().NoError(err)
	s.Assert().NotNil(user)
	s.Assert().EqualValues(1, user.ID)
}

func (s *RepositorySuite) TestCreateWithError() {
	mocket.Catcher.NewMock().
		WithQuery(`INSERT INTO "users"`).
		WithError(gorm.ErrNotImplemented)

	_, err := s.repository.Create("hello", "123456")

	s.Assert().Equal(gorm.ErrNotImplemented, err)
}
