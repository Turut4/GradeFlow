package store

import "gorm.io/gorm"

type ExamReport struct {
	gorm.Model
	ExamID       uint    `json:"exam_id"`
	Exam         Exam    `json:"exam"`
	StudentID    uint    `json:"student_id"`
	TotalScore   float64 `json:"total_score"`
	MaxScore     float64 `json:"max_score"`
	CorrectCount int     `json:"correct_count"`
	WrongCount   int     `json:"wrong_count"`
	Percentage   float64 `json:"percentage"`
	Feedback     string  `gorm:"type:text" json:"feedback"`
}
