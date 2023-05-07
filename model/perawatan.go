package model

import (
	// "database/sql"
	// "mini_project/config"
	"mini_project/config"
	"time"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

type Perawatan struct {
	gorm.Model
	ID          int       `gorm:"primary_key;not null"`
	Name        string    `json:"name" form:"name" gorm:"type:varchar(255)"`
	Keterangan  string    `json:"keterangan" form:"keterangan" gorm:";type:varchar(255)"`
	Harga       float64   `json:"harga" form:"harga" gorm:"type:double"`
	Tanggal     time.Time `json:"tanggalperatawan" form:"tanggal"`
	KambingID   uint      `json:"kambing_id" form:"kambing_id"`
	TransaksiID uint      `json:"transaksi_id" form:"transaksi_id"`
}

var (
// err_perawatan_model error
// DB_perawatan_model  *gorm.DB
)

// func init() {
// 	//connect db
// 	dsn := "root:Mbahbambang123@tcp(localhost:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
// 	DB_perawatan_model, err_perawatan_model = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err_perawatan_model != nil {
// 		panic(err_perawatan_model)
// 	}
// }

func GetAllPerawatan(id_kambing int) (Kambing, error) {
	var kambing Kambing
	err := config.DB.Model(&Kambing{}).Preload("Perawatans").Find(&kambing, id_kambing).Error
	return kambing, err

}

func CreatePerawatan(perawatan Perawatan) (int, int) {
	result := config.DB.Create((&perawatan))
	return int(result.RowsAffected), perawatan.ID
}

func DeletePerawatan(PerawatanID int) int {
	result := config.DB.Delete(&Perawatan{}, PerawatanID)
	return int(result.RowsAffected)
}

func UpdatePerawatan(PerawatanID int, UpdatePerawatan Perawatan) (int, Perawatan) {
	var perawatan Perawatan

	config.DB.Where("id = ?", PerawatanID).First(&perawatan)

	//masih ada kendala di file name dan keterangan jika kosong salah satu maka data yang disimpan nil
	if UpdatePerawatan.Name != "" {
		perawatan.Name = UpdatePerawatan.Name
	}

	if UpdatePerawatan.Keterangan != "" {
		perawatan.Keterangan = UpdatePerawatan.Keterangan
	}

	if UpdatePerawatan.Harga > 0 {
		perawatan.Harga = UpdatePerawatan.Harga
	}

	result := config.DB.Where("id = ?", PerawatanID).Updates(&perawatan)
	return int(result.RowsAffected), perawatan
}
