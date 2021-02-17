package userModel

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id        int
	Email     string
	Firstname string
	Name      string
}
