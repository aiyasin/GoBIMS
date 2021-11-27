package model

import "gorm.io/gorm"

type BookInfo struct {
	gorm.Model
	Id         int     `gorm:"int;not null;columns:id" json:"id"`
	Name       string  `gorm:"type:varchar(50);not null;columns:name" json:"name"`
	Categories string  `gorm:"type:varchar(50);not null;columns:categories" json:"categories"`
	Author     string  `gorm:"type:varchar(50);not null;columns:author" json:"author"`
	Price      float32 `gorm:"type:float32;not null;columns:price" json:"price"`
	Date       string  `gorm:"type:varchar(50);not null;columns:date" json:"date"`
}
