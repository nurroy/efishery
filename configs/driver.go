package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB{
	if DB != nil{
		return DB
	}

	var err error

	DB,err= gorm.Open("sqlite3","./database/efishery.db")
	log.Println("Connected to Database local")

	if err != nil {
		log.Println( "[Configs.ConnectDB] error when connect to database")
		log.Fatal(err)
	} else {
		log.Println( "SUCCES CONNECT TO DATABASE")
	}

	return DB
}
