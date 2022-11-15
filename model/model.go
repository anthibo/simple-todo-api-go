package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	ID        int64  `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Item      string `json:"item"`
	Completed bool   `json:"completed" gorm:"default:false"`
}

type TodoInput struct {
	Item string `json:"item"`
}

var DB *gorm.DB

func ConnectDatabase() {

	dbUrl := "root:password@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Todo{})
	if err != nil {
		return
	}

	DB = database
}
