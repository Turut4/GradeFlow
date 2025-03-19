package store

import (
	"time"

	"gorm.io/gorm"
)

type Grade struct {
	ID        int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	SchoolID  int64          `json:"school_id"`
	School    School         `json:"school"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
