package store

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Test struct {
	ID         int64          `josn:"id" gorm:"primaryKey;autoIncrement"`
	Answers    []string       `json:"answers"`
	Score      int            `json:"score"`
	Gabarito   Gabarito       `json:"gabarito" gorm:"foreignKey:GabaritoID"`
	GabaritoID int64          `json:"gabarito_id"`
	UserID     int64          `json:"user_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type TestStore struct {
	db *gorm.DB
}

func (s *TestStore) Create(ctx context.Context, test *Test) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	return s.db.WithContext(ctx).Select("answers", "gabarito_id", "result").Create(test).Error
}
