package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID uint
	Company   Company `gorm:"foreignKey:CompanyID"`
}

type Company struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("go.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var user User
	var company Company
	err = db.AutoMigrate(&User{}, &Company{})
	if err != nil {
		log.Fatal(err)
	}

	db.Create(&Company{Name: "Developer"})
	db.Save(&company)
	db.Create(&User{Name: "Samuel", CompanyID: company.ID})
	db.Save(&user)
}
