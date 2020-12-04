package database

import (
	"belajar-fiber/model"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectDB() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
}
