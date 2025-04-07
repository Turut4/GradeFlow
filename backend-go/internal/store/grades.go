package store

import (
	"gorm.io/gorm"
)

type Grade struct {
	gorm.Model
	Name     string `gorm:"size:100;not null" json:"name"`
	Students []User `gorm:"many2many:user_grades;" json:"students"`
}
