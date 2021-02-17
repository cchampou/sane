package db

import (
	"../../models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DBORM struct {
	*gorm.DB
}

type DBLayer interface {
	CreateUser(userModel.User) error
}

func Connect() (*DBORM, error) {
	password := os.Getenv("DBPASS")
	if password == "" {
		log.Fatal("Missing password")
	}
	dsn := "host=kandula.db.elephantsql.com user=feasyvsr password=" + os.Getenv("DBPASS") + " dbname=feasyvsr port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to db, migrating...")
	db.AutoMigrate(&userModel.User{})
	log.Print("Migration complete")

	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) CreateUser(user userModel.User) error {
	if &user == nil {
		log.Print("user not defined")
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Print(result.Error)
	}
	log.Print(result.RowsAffected)
	return result.Error
}
