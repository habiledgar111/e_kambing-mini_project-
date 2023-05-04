package model

import (
	// "database/sql"
	// "mini_project/config"

	"mini_project/config"
	"time"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

type Kambing struct {
	gorm.Model
	ID          int         `gorm:"primary_key;not null"`
	Name        string      `json:"name" form:"name" gorm:"type:varchar(255)"`
	TanggalBeli time.Time   `json:"tanggalbeli" form:"tanggalbeli"`
	Status      string      `json:"status" form:"status" gorm:"type:varchar(255)"`
	Harga       float64     `json:"harga" form:"harga" gorm:"type:double"`
	UserID      uint        `json:"user_id" form:"user_id"`
	Perawatans  []Perawatan `json:"perawatans"`
}

var (
// err error
// db  *gorm.DB
)

//	func init() {
//		//connect db
//		dsn := "root:Mbahbambang123@tcp(localhost:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
//		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//		if err != nil {
//			panic(err)
//		}
//	}
func GetAllKambing() ([]Kambing, error) {
	var kambing []Kambing
	// err := db.Model(&Kambing{}).Preload("UserID").Find(&kambing).Error
	err := config.DB.Model(&Kambing{}).Find(&kambing).Error

	return kambing, err
}

func GetKambingByID(id int) (Kambing, error) {
	var kambing Kambing
	err := config.DB.Model(&Kambing{}).Preload("UserID").First(&kambing, id).Error
	return kambing, err
}

func CreateKambingModel(kambing Kambing) int {
	result := config.DB.Create(&kambing)
	return int(result.RowsAffected)
}

func GetAllKambingsfromUser(id int) (User, error) {
	var user User
	err := config.DB.Model(&User{}).Preload("Kambings").Find(&user, id).Error
	return user, err
}
