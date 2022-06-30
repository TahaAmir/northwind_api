package api

import (
	"encoding/json"
	"golang-crud-rest-api/middleware"
	product_querys "golang-crud-rest-api/querys"
	products "golang-crud-rest-api/type"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutesForProducts(router *mux.Router) {

	middleware.EnableCORS(router)

	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		product, err := product_querys.GetProduct()
		if err == nil {
			middleware.RespondWithSuccess(product, w)

		} else {
			middleware.RespondWithError(err, w)

		}
	}).Methods(http.MethodGet)

	//To get by id
	router.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		product, err := product_querys.GetProductById(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(product, w)
		}

	}).Methods(http.MethodGet)

	//To Create
	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {

		var product products.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := product_querys.CreateProduct(product)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	//TO Update
	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {

		var product products.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := product_querys.UpdateProduct(product)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

	//To DELETE
	router.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		err = product_querys.DeleteProduct(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(true, w)
		}

	}).Methods(http.MethodDelete)

}
