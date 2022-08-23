package repository

import (
	"log"
	"os"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var currentDB *sql.DB

// Provide initialisation of the database in the code
func init() {
	godotenv.Load(".env")

	// TODO : Support for information abour database init + handle time if failed

	connStr := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " dbname=" + os.Getenv("DB_NAME") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping DB to test if connection is woring
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Save the db in the package localy so repo can acces it
	currentDB = db
}
