package config

import (
	// "mini_project/model"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Open() error {
	//connect db
	dbUsername := "root"
	dbPassword := "Mbahbambang123"
	dbHost := "miniproject.ck6i8ucy7zfd.us-east-1.rds.amazonaws.com"
	dbName := "miniproject"
	dbPort := "3306"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func OpenLocal() error {
	dsn := "root:Mbahbambang123@tcp(localhost:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	if err != nil {
		return err
	}
	return nil
}
