// db package provides all the database related functionality
package db

import (
	"log"
	"os"
	"strings"

	cfg "app/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDatabase sets up databases for the app
func SetupDatabase() {

	// connecting postgres
	ConnectPostgres()
}

var (
	// PgDB is the postgress connection handle
	PgDB *gorm.DB
)

/////////////////////////
// POSTGRES
/////////////////////////

// ConnectPostgres Returns the Pg DB Instance
func ConnectPostgres() {
	dsn := cfg.GetConfig().Postgres.GetPostgresConnectionInfo()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(strings.Repeat("!", 40))
		log.Println("\tCould Not Establish Postgres DB Connection")
		log.Println(strings.Repeat("!", 40))
		log.Fatal(err)
	}

	log.Println(strings.Repeat("-", 80))
	log.Println("\tConnected To Postgres DB")
	if env := os.Getenv("ENV"); env == "DEV" {
		log.Print(dsn)
	}
	log.Println(strings.Repeat("-", 80))

	PgDB = db
}
