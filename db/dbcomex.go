package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=luca password=test123 dbname=tareas port=5432"
var DB *gorm.DB

func BDconnex() {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println("DB Conectado...")
	}

}
