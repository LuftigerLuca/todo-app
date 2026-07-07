package model

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Title          string `gorm:"size:255;not null" json:"title"`
	OwnerID        uint   `gorm:"not null" json:"owner_id"`
	Owner          User   `gorm:"foreignKey:OwnerID;references:ID" json:"owner"`
	Todos          []Todo `gorm:"foreignKey:ListID" json:"todos"`
	PermittedUsers []User `gorm:"many2many:list_users;" json:"permitted_users"`
}
