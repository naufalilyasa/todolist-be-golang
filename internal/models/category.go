package models

import "time"

type Category struct {
	ID        int      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Color     string    `gorm:"type:varchar(20)" json:"color"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Todos []Todo `gorm:"foreignKey:CategoryID" json:"todos,omitempty"`
}
