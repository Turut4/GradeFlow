package store

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrNotFound          = errors.New("record not found")
	QueryTimeoutDuration = 5 * time.Second
)

type Storage struct {
	Users interface {
		GetByID(ctx context.Context, userID uint) (*User, error)
		GetByEmail(ctx context.Context, email string) (*User, error)
		Create(ctx context.Context, user *User) error
	}
	Roles interface {
		GetByName(ctx context.Context, role string) (*Role, error)
	}
	Exams interface {
		Create(ctx context.Context, exam *Exam) error
		GetByID(ctx context.Context, examID uint) (*Exam, error)
		GetAnswerSheet(ctx context.Context, examID uint) ([]byte, error)
	}
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		Users: &UserStore{db},
		Roles: &RoleStore{db},
		Exams: &ExamStore{db},
	}
}
