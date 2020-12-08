package repository_test

func (s *RepositorySuite) TestCreate() {
	_, err := s.repository.Create("hello", "123456")
	s.Assert().NoError(err)
}
