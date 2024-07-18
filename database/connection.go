package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (err error) {
	dsn := "teste:teste123@tcp(mysql:3306)/testedb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func AutoMigrate(v any) error {
	if DB == nil {
		if err := Connect(); err != nil {
			return err
		}
	}
	return DB.AutoMigrate(v)
}
