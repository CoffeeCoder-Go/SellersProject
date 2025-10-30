package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(){
	var err error

	db,err = gorm.Open(sqlite.Open("./db/development.db"),&gorm.Config{})

	if err != nil {
		log.Printf("Erro no db:%v\n",err.Error())
		return
	}
}

func GetDB() *gorm.DB{
	return  db
}