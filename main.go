package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // blank import
)

var db *gorm.DB

func main() {
	// Open the database
	var err error
	db, err = gorm.Open("postgres", "host=db user=myuser dbname=mydatabase password=example sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
    // Test the database connection
    var servers []Server
    if err := db.Find(&servers).Error; err != nil {
        log.Fatal("Failed to connect to the database:", err)
    } else {
        log.Println("Database connection established successfully")
    }

	r := mux.NewRouter()

	r.HandleFunc("/servers", CreateServerHandler(db)).Methods("POST")
	r.HandleFunc("/servers", GetServersHandler(db)).Methods("GET")
	r.HandleFunc("/servers/{id}", GetServerHandler(db)).Methods("GET")

	log.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", r)
    log.Fatal(http.ListenAndServe(":8000", r))
}
