package store

import (
	"context"

	"gorm.io/gorm"
)

type Role struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int64  `json:"level"`
}

type RoleStore struct {
	db *gorm.DB
}

func (s *RoleStore) GetByName(ctx context.Context, name string) (*Role, error) {
	var role Role

	tx := s.db.WithContext(ctx).Where("name = ?", name).First(&role)

	if tx.Error != nil {
		switch tx.Error {
		case gorm.ErrRecordNotFound:
			return nil, ErrNotFound
		default:
			return nil, tx.Error
		}
	}

	return &role, nil
}
