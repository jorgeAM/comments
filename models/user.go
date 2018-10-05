package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Fullname string `json:"fullname" gorm:"not null"`
	//si viene vacio se va a omitir
	Password string `json:"password,omitempty" gorm:"not null;type:varchar(256)"`
	//no se guardara en db
	ConfirmPassword string `json:"confirmPassword,omitempty" gorm:"-"`
	Avatar          string `json:"avatar"`
	/*User puede tener muchos comentarios*/
	Comments []Comment `json:"comments,omitempty"`
}
