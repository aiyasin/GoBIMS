package model

import (
	"GoBIMS/utils/errmsg"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name       string    `gorm:"type:varchar(50);not null;columns:name" json:"name"`
	Categories string    `gorm:"type:varchar(50);not null;columns:categories" json:"categories"`
	Author     string    `gorm:"type:varchar(50);not null;columns:author" json:"author"`
	Price      float32   `gorm:"type:float;not null;columns:price" json:"price"`
	Date       time.Time `gorm:"type:date;not null;columns:date" json:"date"`
}

// CreateBookInfo 新增图书信息
func CreateBookInfo(data *Book) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// GetBookInfo 查询图书列表
func CheckBookPage(pageSize int, pageNum int) ([]Book, int, int64) {
	var booklist []Book
	var err error
	var total int64
	err = db.Select("id, name, categories, author, price, date, created_at, updated_at, deleted_at").
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Order("id").
		Find(&booklist).Error
	// 单独计数
	db.Model(&booklist).Count(&total)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errmsg.ERROR, 0
	}
	return booklist, errmsg.SUCCESS, total
}

// SearchBookInfo 搜索图书
func SearchBookInfo(name string, categories string, pageSize int, pageNum int) ([]Book, int, int64) {
	var booklist []Book
	var err error
	var total int64
	err = db.Select("id, name, categories, author, price, date, created_at, updated_at, deleted_at").
		Limit(pageSize).
		Offset((pageNum-1)*pageSize).
		Order("id").
		Where("name LIKE ? AND categories LIKE ?", "%"+name+"%", "%"+categories+"%").
		Find(&booklist).
		Count(&total).Error

	if err != nil {
		fmt.Println(err.Error())
		return nil, errmsg.ERROR, 0
	}
	return booklist, errmsg.SUCCESS, total
}

// EditBookInfo 编辑图书
func EditBookInfo(id int, data *Book) int {
	var art Book
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["categories"] = data.Categories
	maps["author"] = data.Author
	maps["price"] = data.Price
	maps["date"] = data.Date

	err = db.Model(&art).Where("id = ? ", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteBookInfo 删除图书
func DeleteBookInfo(id int) int {
	var art Book
	err = db.Where("id = ? ", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
