package store

import (
	"time"

	"gorm.io/gorm"
)

type School struct {
	ID             int64              `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string             `json:"name"`
	Email          string             `json:"email"`
	SubscriptionID int64              `json:"subscription_id"`
	Subscription   SchoolSubscription `json:"subscription" gorm:"foreignKey:SubscriptionID"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	DeletedAt      gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
}

type SchoolSubscription struct {
	ID                 int64            `json:"id" gorm:"primaryKey;autoIncrement"`
	SchoolID           int64            `json:"school_id"`
	SubscriptionPlanID int64            `json:"subscription_plan_id"`
	SubscriptionPlan   SubscriptionPlan `json:"subscription_plan" gorm:"foreignKey:SubscriptionPlanID"`
	StartDate          time.Time        `json:"start_date"`
	EndDate            time.Time        `json:"end_date"`
	IsActive           bool             `json:"is_active"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
	DeletedAt          gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
}
