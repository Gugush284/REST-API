package teststore

import (
	model_user "github.com/Gugush284/Go-server.git/internal/app/model/user"
	"github.com/Gugush284/Go-server.git/internal/app/store"
)

// Store for clients ...
type TestStore struct {
	userRepository *UserRepository
}

// New Store ...
func New() *TestStore {
	return &TestStore{}
}

// Access for User
func (s *TestStore) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model_user.User),
	}

	return s.userRepository
}
