package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Client *sqlx.Conn

// ConnectDatabase is used to connect the MongoDB database
func ConnectDatabase() {
	log.Println("Database connecting...")

	// Connect to MongoDB
	Client, err := sqlx.Connect("postgres", "user=root password=secret dbname=root sslmode=disable")
	if err != nil {
		log.Println("Database Connection Failed!")
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connected.")
}
