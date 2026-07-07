package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Title string `gorm:"size:100;not null;uniqueIndex" json:"title"`
	Todos []Todo `gorm:"many2many:todo_tags;" json:"-"`
}
