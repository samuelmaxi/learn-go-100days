package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func authenticateUser(db *gorm.DB, username, password string) bool {
	var user User
	result := db.Where("user_name = ?", username).First(&user)
	if result.Error != nil {
		fmt.Println("User not found")
		return false
	}

	if user.Password != password {
		fmt.Println("Incorrect password")
		return false
	}

	fmt.Println("Login")
	return true
}

func main() {
	db, err := gorm.Open(sqlite.Open("go.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var user User
	db.AutoMigrate(&User{})
	user = User{UserName: "Samuel", Password: "1234"}
	db.Create(&user)
	db.Save(&user)

	authenticateUser(db, "Samuel", "1234")
	authenticateUser(db, "Sam", "4321")

}
