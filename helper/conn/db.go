package conn

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var client *sqlx.DB

func New() *sqlx.DB {
	if client == nil {
		connectDatabase()
	}
	return client
}

// ConnectDatabase is used to connect the postgres database
func connectDatabase() {
	var err error
	log.Println("Database connecting...")

	// Connect to postgres
	client, err = sqlx.Connect("postgres", "user=root password=secret dbname=root sslmode=disable")
	if err != nil {
		log.Println("Database Connection Failed!")
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connected.")
}
