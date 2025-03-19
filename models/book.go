package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" gorm:"type:varchar(100);not null"`
	Author string `json:"author" gorm:"type:varchar(100);not null"`
	Price  int    `json:"price" gorm:"not null"`
}
