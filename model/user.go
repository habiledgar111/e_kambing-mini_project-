package model

import (
	// "database/sql"
	// "mini_project/config"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primary_key;not null"`
	Email    string `gorm:";type:varchar(255)unique;not null"`
	Password string `gorm:"notnull"`
	Name     string `gorm:"type:varchar(255)"`
}
type UserLogin struct {
	email    string
	password string
}
type UserMock struct {
	gorm.Model
	Email    string
	Password string
	Name     string
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "users"
}

var (
	user *User
)

func init() {
	//connect db
	// dsn := "root:Mbahbambang123@tcp(localhost:3306)/sec21orm?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// DB = db
	// DB.AutoMigrate(user)
}

// func GetUserFromEmail(userlogin UserLogin) (interface{}, interface{}) {
// 	if err := DB.Where("email = @email", sql.Named("email", userlogin.email)).First(&user).Error; err != nil {
// 		return nil,err
// 	}
// 	return user,nil
// }
