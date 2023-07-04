package main

import (
	"log"
	"fmt"
	"os"
	"net/http"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // blank import
)

var db *gorm.DB

func main() {
	// Open the database
	var err error
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbUser, dbName, dbPassword)

	db, err = gorm.Open("postgres", connectionString)
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

	handler := cors.Default().Handler(r)

	log.Println("Starting server on port 8000")
    log.Fatal(http.ListenAndServe(":8000", handler))
}
