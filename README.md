# northwind_api

Clone this repo

Install the dependencies with:

go get github.com/go-sql-driver/mysql
go get github.com/gorilla/mux
go get github.com/joho/godotenv


Create a database in MySQL

Create a file called .env 

Configure credentials in .env

In the database create the tables according to northwind.sql

Compile with go run .

Now you can test the API with postman, it will be on localhost:8000