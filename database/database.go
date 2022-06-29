package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")
var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("USER"),
		os.Getenv("PASS"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DB_NAME"))
)

const AllowedCORSDomain = "http://localhost"

var DB *sql.DB

//function to connect MySql to Go
func GetDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}
