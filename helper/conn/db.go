package conn

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lukamindo/pet-reminder/app/constant"
)

var client *sqlx.DB

type Config struct {
	Host    string
	User    string
	Pass    string
	Port    string
	DBName  string
	SSLMode string
}

func config() Config {
	return Config{
		Host:    os.Getenv(constant.DBHostKey),
		User:    os.Getenv(constant.DBUserKey),
		Pass:    os.Getenv(constant.DBPassKey),
		Port:    os.Getenv(constant.DBPortKey),
		DBName:  os.Getenv(constant.DBDBNameKey),
		SSLMode: os.Getenv(constant.DBSSLMode),
	}
}

func New() *sqlx.DB {
	if client == nil {
		config := config()
		config.connectDatabase()
	}
	return client
}

// ConnectDatabase is used to connect the postgres database
func (c Config) connectDatabase() {
	var err error
	log.Println("Database connecting...")
	// Connect to postgres
	client, err = sqlx.Connect("postgres", c.dbConnectionString())
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

func (c Config) dbConnectionString() string {
	conn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		c.Host,
		c.User,
		c.Pass,
		c.DBName,
		c.Port,
		c.SSLMode,
	)
	return conn
}
