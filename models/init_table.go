package models

import "github.com/jinzhu/gorm"

func InitTable(db *gorm.DB) {

	db.AutoMigrate(&Auth{})
	db.AutoMigrate(&User{})
}