package orm

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `gorm:"primary_key;-"`
	Username string `gorm:"primary_key"`
	Password string
	Fname string
	Lname string
	Email string
}