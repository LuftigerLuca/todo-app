package model

import (
	"time"

	"gorm.io/gorm"
)

// @Enum pending in_progress completed
type TodoStatus string

const (
	StatusPending    TodoStatus = "pending"
	StatusInProgress TodoStatus = "in_progress"
	StatusCompleted  TodoStatus = "completed"
)

type Todo struct {
	gorm.Model
	Title       string     `gorm:"size:255;not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	Status      TodoStatus `gorm:"size:20;default:pending" json:"status"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	ListID      uint       `gorm:"not null;index" json:"list_id"`
	Tags        []Tag      `gorm:"many2many:todo_tags;" json:"tags"`
}
