package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-crud-gorm-mux-sql/models"
	"log"
	"net/http"
	"strconv"
)

var err error

type response struct {
	ID      uint        `json:"id,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var albums []models.Album
	models.AlbumIndex(&albums)

	if err := json.NewEncoder(w).Encode(albums); err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var album models.Album

	params := mux.Vars(r)
	id := params["id"]

	album = models.AlbumByID(id)

	if err := json.NewEncoder(w).Encode(album); err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var album models.Album

	if err = json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, "Error decoding request body", http.StatusInternalServerError)
		log.Fatal(err)
	}

	models.AlbumCreate(&album)

	res := response{
		ID:      album.ID,
		Message: "User created successfully",
		Data:    album,
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var album models.Album

	params := mux.Vars(r)
	id := params["id"]
	id64, _ := strconv.ParseUint(id, 0, 64)
	idToUpdate := uint(id64)
	album.ID = idToUpdate

	if err = json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, "Error decoding request body", http.StatusInternalServerError)
		log.Fatal(err)
	}

	models.AlbumUpdate(album)

	res := response{
		ID:      album.ID,
		Message: "User updated successfully",
		Data:    album,
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	id := params["id"]
	id64, _ := strconv.ParseUint(id, 0, 64)
	idToDelete := uint(id64)

	models.AlbumDelete(id)

	res := response{
		ID:      idToDelete,
		Message: "User deleted successfully",
	}

	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		log.Fatal(err)
	}
}
