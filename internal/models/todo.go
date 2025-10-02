package models

import "time"

type Priority string

const (
	PriorityHigh   Priority = "high"
	PriorityMedium Priority = "medium"
	PriorityLow    Priority = "low"
)

type Todo struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Priority    Priority  `gorm:"type:priority_type;default:'medium'" json:"priority"`
	IsCompleted bool      `gorm:"default:false" json:"is_completed"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	CategoryID *int      `gorm:"index" json:"category_id,omitempty"`
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}
