package main

import (
	"log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func CreateServerHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var server Server
		if err := json.NewDecoder(r.Body).Decode(&server); err != nil {
			log.Println(&server)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.Create(&server).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(&server)
	}
}

func GetServersHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var servers []Server

		if err := db.Find(&servers).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(&servers)
	}
}

func GetServerHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var server Server
		params := mux.Vars(r)

		if err := db.First(&server, params["id"]).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(&server)
	}
}
