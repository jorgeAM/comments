package models

import "github.com/jinzhu/gorm"

// Vote model
type Vote struct {
	gorm.Model
	CommentID uint `json:"comentId" gorm:"not null"`
	UserID    uint `json:"userId" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
