package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// UserAccount is database table which store user account infomation
type UserAccount struct {
	gorm.Model
	ID       string
	Account  string `form:"account" json:"account" binding:"required" gorm:"type:varchar(100);unique_index"`
	Password string `form:"password" json:"password" binding:"required" gorm:"type:varchar(100);not null"`
}

func dbSearchUser(user UserAccount) bool {
	var usert UserAccount
	if db.Where("account = ? AND password = ?", user.Account, user.Password).First(&usert).RecordNotFound() {
		return false
	}

	return true
}

func initDB() {
	var err error
	if db, err = gorm.Open("mysql", "root:Nozuonodie@/booksalon?charset=utf8&parseTime=True&loc=Local"); err != nil {
		println("mysql DB open error:", err)
		os.Exit(0)
	}
	user := UserAccount{Account: "jelech", Password: "Nozuonodie"}
	if err := db.AutoMigrate(&UserAccount{}).Error; err != nil {
		println(err)
		os.Exit(0)
	}
	db.Create(&user)
}
