package main

import (
	"github.com/gorilla/mux"
	"go-crud-gorm-mux-sql/controllers"
	"go-crud-gorm-mux-sql/models"
	"log"
	"net/http"
)

func init() {
	models.Setup()
}

func main() {

	log.Println("Starting server at http://127.0.0.1:8080/")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/albums/{id}", controllers.GetAlbum).Methods("GET")
	router.HandleFunc("/api/albums", controllers.GetAlbums).Methods("GET")
	router.HandleFunc("/api/albums", controllers.CreateAlbum).Methods("POST")
	router.HandleFunc("/api/albums/{id}", controllers.UpdateAlbum).Methods("PUT")
	router.HandleFunc("/api/albums/{id}", controllers.DeleteAlbum).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
