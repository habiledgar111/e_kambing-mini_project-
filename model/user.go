package model

import (
	// "database/sql"
	// "mini_project/config"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

type User struct {
	gorm.Model
	ID         int         `gorm:"primary_key;not null"`
	Email      string      `json:"email" form:"email" gorm:";type:varchar(255)unique;not null"`
	Password   string      `json:"password" form:"password" gorm:"notnull"`
	Name       string      `json:"name" form:"name" gorm:"type:varchar(255)"`
	Kambings   []Kambing   `json:"kambings"`
	Transaksis []Transaksi `json:"traksaksis"`
}
