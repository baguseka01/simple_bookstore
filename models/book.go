package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          int    `gorm:"column:id;primaryKey;autoIncrement"`
	ISBN        string `gorm:"column:isbn"`
	Title       string `gorm:"column:title"`
	Author      string `gorm:"column:author"`
	Price       int64  `gorm:"column:price"`
}
