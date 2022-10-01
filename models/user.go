package models

import "time"

type User struct {
	ID        int `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time
	UserID    string `gorm:"size:20;not null;" json:"userid"`
	Username  string `gorm:"size:15;not null" json:"username"`
	Password  string `gorm:"size:255;not null" json:"password"`
	Role      string `gorm:"size:10" json:"role"`
	Token 	  string `sql:"-" json:"token"`
}


// TableName ..
func (s User) TableName() string {
	return "tb_user"
}

type LoginResponse struct {
	Token 	  string `json:"token"`
}

type RegisterReponse struct {
	Password string `json:"password"`
}

type ValidateResponse struct {
	UserID    string `json:"userid"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Exp 	  int64  `json:"exp"`
}