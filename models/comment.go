package models

import "github.com/jinzhu/gorm"

// Comment model
type Comment struct {
	gorm.Model
	UserID   uint   `json:"userId"`
	ParentID uint   `json:"parentId"`
	Votes    uint32 `json:"votes"`
	Content  string `json:"content"`
	//no se guardara en db
	HasVote uint8 `json:"hasVote" gorm:"-"`
	/*
	  User     []User    `json:"user,omitempty"`
	*/
	Children []Comment `json:"comments,omitempty"`
}
