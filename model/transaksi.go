package model

import (
	// "database/sql"
	// "mini_project/config"
	"mini_project/config"
	"strconv"
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

func DeleteTransaksi(id int) int {
	result := config.DB.Delete(&Transaksi{}, id)
	return int(result.RowsAffected)
}

func DeleteAllTransaksifromKambing(id int) int {
	result := config.DB.Where("kambing_id = ?", id).Delete(&Transaksi{})
	return int(result.RowsAffected)
}

func DeleteAllTransaksifromPerawatan(id int) int {
	result := config.DB.Where("perawatan_id = ?", id).Delete(&Transaksi{})
	return int(result.RowsAffected)
}

func UpdateTransaksi(id int, updatetransaksi Transaksi) (int, Transaksi) {
	var transaksi Transaksi
	config.DB.Where("id = ?", id).First(&transaksi)

	//masih ada kendala di file name dan keterangan jika kosong salah satu maka data yang disimpan nil
	if updatetransaksi.Name != "" {
		transaksi.Name = updatetransaksi.Name
	}

	if updatetransaksi.Keterangan != "" {
		transaksi.Keterangan = updatetransaksi.Keterangan
	}

	if updatetransaksi.Harga > 0 {
		transaksi.Harga = updatetransaksi.Harga
	}

	result := config.DB.Where("id = ?", id).Updates(&transaksi)
	return int(result.RowsAffected), transaksi
}

func UpdateTransaksifromKambing(id int, updatetransaksi Transaksi) (int, Transaksi) {
	var transaksi Transaksi
	config.DB.Where("kambing_id = ?", id).First(&transaksi)

	//masih ada kendala di file name dan keterangan jika kosong salah satu maka data yang disimpan nil
	if updatetransaksi.Name != "" {
		transaksi.Name = updatetransaksi.Name
	}

	transaksi.Keterangan = "membeli kambing - " + strconv.Itoa(id)

	if updatetransaksi.Harga > 0 {
		transaksi.Harga = updatetransaksi.Harga
	}

	result := config.DB.Where("kambing_id = ?", id).Updates(&transaksi)
	return int(result.RowsAffected), transaksi
}

func UpdateTransaksifromPerawatan(id int, updatetransaksi Transaksi) (int, Transaksi) {
	var transaksi Transaksi
	config.DB.Where("perawatan_id = ?", id).First(&transaksi)

	//masih ada kendala di file name dan keterangan jika kosong salah satu maka data yang disimpan nil
	if updatetransaksi.Name != "" {
		transaksi.Name = updatetransaksi.Name
	}

	if updatetransaksi.Keterangan != "" {
		transaksi.Keterangan = updatetransaksi.Keterangan
	}

	if updatetransaksi.Harga > 0 {
		transaksi.Harga = updatetransaksi.Harga
	}

	result := config.DB.Where("perawatan_id = ?", id).Updates(&transaksi)
	return int(result.RowsAffected), transaksi
}
