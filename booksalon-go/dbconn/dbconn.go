package dbconn

import (
	"os"

	"github.com/jinzhu/gorm"
	// this will user into gorm to control the db
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// LoginInfo is login json message that must need
type LoginInfo struct {
	Name     string
	Account  string
	Password string
}

// UserAccount is database table which store user account infomation
type UserAccount struct {
	gorm.Model
	Account  string `form:"account" json:"account" binding:"required" gorm:"type:varchar(100);unique_index"`
	Password string `form:"password" json:"password" binding:"required" gorm:"type:varchar(100);not null"`
	User     User   `binding:"-"`
}

// User is student table in the mysql
type User struct {
	gorm.Model
	// userCanShow
	Name      string `form:"name" json:"name" binding:"required" gorm:"type:varchar(100);not null"`
	Teams     []Team `gorm:"many2many:user_teams"`
	AccountID uint   `binding:"-"`

	// UserAccount UserAccount // this will get user password!!! which must can't be shown
}

// Team is a student group
type Team struct {
	gorm.Model
	Topic    string `form:"topic" json:"topic" binding:"required" gorm:"type:varchar(100);not null"`
	LeaderID string `form:"leaderid" json:"leaderid" binding:"required" gorm:"type:varchar(100);not null"`
	Users    []User `gorm:"many2many:user_teams"`
}

// NewDBConn will create a gorm db connect, and return it
func NewDBConn() *gorm.DB {
	var err error
	if db, err = gorm.Open("mysql", "root:Nozuonodie@/booksalon?charset=utf8&parseTime=True&loc=Local"); err != nil {
		println("mysql DB open error:", err.Error())
		os.Exit(0)
	}

	db.LogMode(true)
	db.AutoMigrate(&UserAccount{}) // 更新表结构
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Team{})
	db.Model(&UserAccount{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	// db.Create(&UserAccount{Account: "jelech1", Password: "123"})
	// db.Create(&UserAccount{Account: "jelech2", Password: "123"})
	// db.Create(&UserAccount{Account: "jelech3", Password: "123"})

	return db
}
