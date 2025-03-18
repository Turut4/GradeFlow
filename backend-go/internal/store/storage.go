package store

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("record not found")
)

type Storage struct {
	Users interface {
		GetByID(ctx context.Context, userID int64) (*User, error)
		GetByEmail(ctx context.Context, email string) (*User, error)
		Create(ctx context.Context, user *User) error
	}
	Roles interface {
		GetByName(ctx context.Context, role string) (*Role, error)
	}
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		Users: &UserStore{db},
		Roles: &RoleStore{db},
	}
}
