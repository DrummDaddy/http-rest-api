package teststore

import (
	"github.com/gopherschool/http-rest-api/cmd/apiserver/internal/app/model"
	"github.com/gopherschool/http-rest-api/cmd/apiserver/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}

}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository

}
