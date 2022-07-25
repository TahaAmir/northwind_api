package api

import (
	"encoding/json"
	"golang-crud-rest-api/middleware"
	"golang-crud-rest-api/querys"
	types "golang-crud-rest-api/type"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SetupRoutesForProducts(router *mux.Router) {

	middleware.EnableCORS(router)

	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {

		count, _ := strconv.Atoi(r.FormValue("count"))
		start, _ := strconv.Atoi(r.FormValue("start"))

		if count > 10 || count < 1 {
			count = 10
		}
		if start < 0 {
			start = 0
		}

		product, err := querys.GetProduct(start, count)
		if err == nil {
			middleware.RespondWithJSON(w, http.StatusOK, product)

		} else {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())

		}
	}).Methods(http.MethodGet)

	//To get by id
	router.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		product, err := querys.GetProductById(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusNotFound, "Product not found")
			return
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, product)
		}

	}).Methods(http.MethodGet)

	//To Create
	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {

		var product types.Product
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		json.Unmarshal([]byte(body), &product)

		res, err := querys.CreateProduct(product)
		if err != nil {

			middleware.RespondWithError(w, http.StatusConflict, err.Error())
			return
		}
		middleware.RespondWithJSON(w, http.StatusCreated, res)
	}).Methods(http.MethodPost)

	//TO Update
	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {

		var product types.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		} else {
			err := querys.UpdateProduct(product)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				middleware.RespondWithJSON(w, http.StatusOK, product)
			}
		}
	}).Methods(http.MethodPut)

	//To DELETE
	router.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		err = querys.DeleteProduct(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
		}

	}).Methods(http.MethodDelete)

}
