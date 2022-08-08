package models

import (
	"fmt"
	"gorm.io/gorm"
)

func Setup() {

}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&Album{})
	if err != nil {
		fmt.Printf("Database automigrate error %v", err)
	}
}
