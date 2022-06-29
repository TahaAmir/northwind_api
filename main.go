package main

import (
	product_api "golang-crud-rest-api/api"
	"golang-crud-rest-api/databaseconnection"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	var err error
	databaseconnection.DB, err = databaseconnection.GetDB()
	if err != nil {
		log.Println("Error with database" + err.Error())
		return
	} else {
		err = databaseconnection.DB.Ping()

		if err != nil {
			log.Println("Error conneting to the database. The error is : " + err.Error())
			return
		}

	}

	router := mux.NewRouter()
	product_api.SetupRoutesForProducts(router)

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
