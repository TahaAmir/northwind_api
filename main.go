package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//Defining the connection string to database
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
func getDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}

func main() {

	fmt.Println(ConnectionString)
	var err error
	DB, err = getDB()
	if err != nil {
		log.Println("Error with database" + err.Error())
		return
	} else {
		err = DB.Ping()

		if err != nil {
			log.Println("Error conneting to the database. The error is : " + err.Error())
			return
		}

	}

	router := mux.NewRouter()
	SetupRoutesForProducts(router)

	port := ":8000"

	server := &http.Server{
		Handler: router,
		Addr:    port,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())

}
