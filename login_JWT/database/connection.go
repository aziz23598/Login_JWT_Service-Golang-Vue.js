package database

import (
	"login_JWT/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	connection, err := gorm.Open(mysql.Open("root:@/go"), &gorm.Config{})

	if err != nil {
	panic("could not connect to databases")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
} 