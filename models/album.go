package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}
