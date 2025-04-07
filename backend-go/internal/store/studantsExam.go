package store

import (
	"time"

	"gorm.io/gorm"
)

type StudentExam struct {
	gorm.Model
	ExamID      uint      `json:"exam_id"`
	StudentID   uint      `json:"student_id"`
	Answers     string    `gorm:"type:json" json:"answers"`
	Score       float64   `json:"score"`
	Correct     int       `json:"correct"`
	Total       int       `json:"total"`
	SubmittedAt time.Time `json:"submitted_at"`
}
