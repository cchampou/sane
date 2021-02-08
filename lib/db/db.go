package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect() {

	dsn := "host=kandula.db.elephantsql.com user=feasyvsr password=" + os.Getenv("DBPASS") + " dbname=feasyvsr port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to db")
}
