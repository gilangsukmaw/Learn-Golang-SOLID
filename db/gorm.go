package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDatabase() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/solid_exercise?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err --> ", err)
		return nil, err
	}

	return db, nil
}
