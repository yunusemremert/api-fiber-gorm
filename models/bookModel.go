package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
