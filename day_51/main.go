package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("go.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Samuel",
		Age:  21,
	}

	db.AutoMigrate(&User{})
	result := db.Create(&user)
	fmt.Println(result)

	db.First(&user)
}
