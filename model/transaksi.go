package model

import (
	// "database/sql"
	// "mini_project/config"
	"time"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

type Transaksi struct {
	gorm.Model
	ID         int    `gorm:"primary_key;not null"`
	Keterangan string `gorm:";type:varchar(255)unique;not null"`
	Tanggal    time.Time
	Name       string `gorm:"type:varchar(255)"`
	IDKambing  int
}
