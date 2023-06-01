package models

type User struct {
	ID   uint64 `gorm:"size:5;not null;uniqueIndex;primary_key;autoIncrement:true;"`
	Name string `gorm:"size:255;not null;"`
}
