package models

import (
	"time"
)

// Task represents a todo item in the database
type Task struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string    `gorm:"type:varchar(1000);not null" json:"content"`
	Completed bool      `gorm:"default:false;not null" json:"completed"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName specifies the table name for the Task model
func (Task) TableName() string {
	return "tasks"
}
