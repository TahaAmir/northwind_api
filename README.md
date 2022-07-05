# northwind_api

Clone this repo

# Install the dependencies with:

- go get github.com/go-sql-driver/mysql
- go get github.com/gorilla/mux
- go get github.com/joho/godotenv

 # API's Implemented:

- GET     /api
- GET     /api/{id} 
- POST    /api
- PUT     /api 
- DELETE  /api/{id}  

# Databae Tables 

# API Reference

- Get  /table name or any table u want to get. It will get all the data of the table requested.
- Get  /table name/{id}  It will get all the data of the table by Id that is given.
- POST /table name  will create a new column in the table.
- PUT  /table name  will update the table. 
- Delete /table name/{id}  Will delete the coloumn in  table by a given Id. 

# Project is created with:

- Golang
- gorilla/mux
- joho/godotenv
- MYSQL

# To start 
- Create a database in MySQL

- Create a file called .env 

- Configure credentials in .env

- In the database create the tables according to northwind.sql

- Compile with go run .

- Now you can test the API with postman, it will be on localhost:8000



