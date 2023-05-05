package model

import (
	// "database/sql"
	// "mini_project/config"
	"mini_project/config"
	"time"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

type Transaksi struct {
	gorm.Model
	ID          int     `gorm:"primary_key;not null"`
	Name        string  `gorm:"type:varchar(255)"`
	Keterangan  string  `gorm:";type:varchar(255)"`
	KambingID   uint    `json:"kambing_id" form:"kambing_id"`
	UserID      int     `json:"user_id" form:"user_id"`
	PerawatanID int     `json:"perawatan_id" form:"perawatan_id"`
	Harga       float64 `json:"harga" from:"harga" gorm:"type:double"`
	Tanggal     time.Time
}

func GetAllTransaksifromUser(userID int) (User, error) {
	var user User
	err := config.DB.Model(&User{}).Preload("Transaksis").Find(&user, userID).Error
	return user, err
}

func CreateTransaksifromUser(transaksi Transaksi) int {
	result := config.DB.Omit("KambingID", "PerawatanID").Create(&transaksi)
	return int(result.RowsAffected)
}

func CreateTransaksifromKambing(transaksi Transaksi) int {
	result := config.DB.Omit("UserID", "PerawatanID").Create(&transaksi)
	return int(result.RowsAffected)
}

func CreateTransaksifromPerawatan(Transaksi Transaksi) int {
	result := config.DB.Omit("KambingID", "UserID").Create(&Transaksi)
	return int(result.RowsAffected)
}
