package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func AlbumIndex(albums *[]Album) {
	DB.Find(&albums)
	return
}

func AlbumByID(id string) (album Album) {
	DB.First(&album, id)
	return
}

func AlbumCreate(album *Album) {
	DB.Create(&album)
	return
}

func AlbumUpdate(album Album) {
	DB.Save(&album)
	return
}

func AlbumDelete(id string) (album Album) {
	DB.Delete(&album, id)
	return
}
