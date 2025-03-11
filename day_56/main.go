package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("go.db"), &gorm.Config{})
	if err == nil {
		log.Fatal(err)
	}

	var user User
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "Samuel", Age: 21})
	db.Save(&user)

}
