package store

import "gorm.io/gorm"

type Exam struct {
	gorm.Model
	Title       string `gorm:"size:200;not null" json:"title"`
	Subject     string `gorm:"size:100" json:"subject"`
	GradeLevel  string `gorm:"size:50" json:"grade_level"`
	AnswerSheet string `gorm:"type:json" json:"answer_sheet"`
	Weights     string `gorm:"type:json" json:"weights"`
}
