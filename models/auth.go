package models

import "github.com/jinzhu/gorm"

type Auth struct {
	gorm.Model
	UserID   string `gorm:"size:20;not null;" json:"userid"`
	Username string `gorm:"size:15;not null;" json:"username"`
	Role     string `json:"role"`
	Exp 	 int64  `json:"exp"`
}
