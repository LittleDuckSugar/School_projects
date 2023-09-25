package main

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"

	"fmt"

	_ "github.com/lib/pq"
)

// Env holds database connection to Postgres
var currentDB *sql.DB

// ConnectDB tries to connect DB and on succcesful it returns
// DB connection string and nil error, otherwise return empty DB and the corresponding error.
func init() {
	godotenv.Load(".env")

	psqlInfo := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to DB!")

	currentDB = db
}
