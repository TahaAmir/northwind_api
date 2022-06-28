package databaseconnection

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")
var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
)

const AllowedCORSDomain = "http://localhost"

var DB *sql.DB

//function to connect MySql to Go
func GetDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}
