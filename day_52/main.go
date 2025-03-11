package main

import (
	"encoding/json"
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

	db.AutoMigrate(&User{})

	user := &User{}
	result := db.First(&user)

	bytes, _ := json.Marshal(user)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Print(string(bytes))

}
