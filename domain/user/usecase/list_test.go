package usecase_test

import (
	"errors"
	"go-boilerplate/model"
)

func (s Suite) TestList() {
	expectedUsers := []model.User{{1, "user@mail.com", "randomPassword"}}
	s.repository.On("FindAll").Return(expectedUsers, nil).Once()

	users, err := s.usecase.List()

	s.Nil(err)
	s.EqualValues(expectedUsers, users)
}

func (s Suite) TestListWithError() {
	var expectedUsers []model.User
	expectedError := errors.New("not implemented")
	s.repository.On("FindAll").Return(expectedUsers, expectedError).Once()

	users, err := s.usecase.List()

	s.NotNil(err)
	s.EqualValues(expectedError, err)
	s.EqualValues(expectedUsers, users)
}
