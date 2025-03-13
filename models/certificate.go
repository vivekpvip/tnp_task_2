package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Certificate struct {
	ID      int    `json:"id" gorm:"primarykey"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Owner   string `json:"owner"`
	Date    int    `json:"date"`
}

// InitDB initializes the database connection.
func InitDB() {
	var err error
	dsn := "user=your_user password=your_password dbname=certificate_db host=localhost port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB.AutoMigrate(&Certificate{})
}
