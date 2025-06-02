package store

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:100;unique;not null"                                         json:"username"`
	Email    string `gorm:"size:100;uniqueIndex;not null"                                    json:"email"`
	Password string `gorm:"size:255;not null"                                                json:"-"`
	RoleID   int64  `                                                                        json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
}

type UserStore struct {
	db *gorm.DB
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (u *User) ComparePassword(pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
}

func (s *UserStore) GetByID(ctx context.Context, userID uint) (*User, error) {
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

	return s.db.WithContext(ctx).
		Select("username", "email", "password", "role_id").
		Create(user).
		Error
}

func (s *UserStore) GetByEmail(
	ctx context.Context,
	email string,
) (*User, error) {
	var user User

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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
