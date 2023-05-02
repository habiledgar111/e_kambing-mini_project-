package model

import (
	// "database/sql"
	// "mini_project/config"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

type User struct {
	gorm.Model
	ID       int       `gorm:"primary_key;not null"`
	Email    string    `gorm:";type:varchar(255)unique;not null"`
	Password string    `gorm:"notnull"`
	Name     string    `gorm:"type:varchar(255)"`
	Kambings []Kambing `gorm:"Foreignkey:UserID;association_foreignkey:ID;"`
}
