package store

import (
	"gorm.io/gorm"
)

type Grade struct {
	gorm.Model
	Name     string `gorm:"size:100;not null" json:"name"`
	Students string `gorm:"size:100;not null" json:"students"`
}
