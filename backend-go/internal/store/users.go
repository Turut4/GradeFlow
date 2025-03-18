package store

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string         `gorm:"size:100;unique;not null" json:"username"`
	Email     string         `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"size:72;not null" json:"-"`
	RoleID    int64          `json:"role_id"`
	Role      Role           `gorm:"foreignKey:RoleID" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserStore struct {
	db *gorm.DB
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *UserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	var user User
	tx := s.db.Preload("Role").First(&user, userID)
	if tx.Error != nil {
		switch tx.Error {
		case gorm.ErrRecordNotFound:
			return nil, ErrNotFound
		default:
			return nil, tx.Error
		}
	}
	return &user, nil
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		return err
	}
	user.Password = hashedPassword

	var role Role
	if err := s.db.WithContext(ctx).Where("name = ?", user.Role.Name).First(&role).Error; err != nil {
		return fmt.Errorf("role not found: %w", err)
	}
	user.RoleID = role.ID

	return s.db.WithContext(ctx).Select("username", "email", "password", "role_id").Create(user).Error
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	if err := s.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
