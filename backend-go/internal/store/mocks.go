package store

import (
	"context"

	"gorm.io/gorm"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct{}

func (s *MockUserStore) GetByID(
	ctx context.Context,
	userID uint,
) (*User, error) {
	return &User{Model: gorm.Model{ID: uint(userID)}}, nil
}

func (s *MockUserStore) GetByEmail(
	ctx context.Context,
	email string,
) (*User, error) {
	return &User{
		Model:    gorm.Model{ID: 1},
		Email:    email,
		Password: "hashed-test-password",
	}, nil
}

func (s *MockUserStore) Create(
	ctx context.Context,
	user *User,
) error {
	return nil
}
