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
	TransaksiID uint        `json:"transaksi_id" form:"transaksi_id"`
}

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

func CreateKambingModel(kambing Kambing) (int, int) {
	result := config.DB.Create(&kambing)
	return int(result.RowsAffected), kambing.ID
}

func GetAllKambingsfromUser(id int) (User, error) {
	var user User
	err := config.DB.Model(&User{}).Preload("Kambings").Find(&user, id).Error
	return user, err
}

func UpdateKambing(id int, updatekambing Kambing) (int, Kambing) {
	var kambing Kambing
	var tempkambing Kambing

	config.DB.Where("id = ?", id).First(&kambing)
	config.DB.Where("id = ?", id).First(&tempkambing)

	//masih ada kendala di file name dan keterangan jika kosong salah satu maka data yang disimpan nil
	if updatekambing.Name != "" {
		kambing.Name = updatekambing.Name
	} else {
		kambing.Name = tempkambing.Name
	}

	if updatekambing.Status != "" {
		kambing.Status = updatekambing.Status
	} else {
		kambing.Status = tempkambing.Status
	}

	if updatekambing.Harga > 0 {
		kambing.Harga = updatekambing.Harga
	} else {
		kambing.Harga = tempkambing.Harga
	}

	result := config.DB.Where("id = ?", id).Updates(&kambing)
	return int(result.RowsAffected), kambing
}

func DeleteKambing(id int) int {
	result := config.DB.Delete(&Kambing{}, id)
	return int(result.RowsAffected)
}
