package store

import "context"

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct{}

func (s *MockUserStore) GetByID(
	ctx context.Context,
	userID int64,
) (*User, error) {
	return &User{}, nil
}

func (s *MockUserStore) GetByEmail(
	ctx context.Context,
	email string,
) (*User, error) {
	return &User{}, nil
}

func (s *MockUserStore) Create(
	ctx context.Context,
	user *User,
) error {
	return nil
}
