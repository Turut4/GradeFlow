package store

import (
	"time"

	"gorm.io/gorm"
)

type Gabarito struct {
	ID             int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string         `json:"name"`
	CorrectAnswers []string       `json:"correct_answers" gorm:"type:text[]"`
	Grade          Grade          `json:"grade" gorm:"foreignKey:GradeID"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
