package model

import (
	// "database/sql"
	// "mini_project/config"

	"time"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

type Kambing struct {
	gorm.Model
	ID          int    `gorm:"primary_key;not null"`
	Name        string `gorm:"type:varchar(255)"`
	TanggalBeli time.Time
	Status      string  `gorm:"type:varchar(255)"`
	Harga       float64 `gorm:"type:double"`
	UserID      int
}

var (
	err error
	db  *gorm.DB
)

func init() {
	//connect db
	dsn := "root:Mbahbambang123@tcp(localhost:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
func GetAllKambing() ([]Kambing, error) {
	var kambing []Kambing
	err := db.Model(&Kambing{}).Preload("UserID").Find(&kambing).Error
	return kambing, err
}

func GetKambingByID(id int) (Kambing, error) {
	var kambing Kambing
	err := db.Model(&Kambing{}).Preload("UserID").First(&kambing, id).Error
	return kambing, err
}

func CreateKambingModel(kambing Kambing) int {
	result := db.Create(&kambing)
	return int(result.RowsAffected)
}
