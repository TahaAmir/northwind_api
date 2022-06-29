package product_api

import (
	"encoding/json"
	"golang-crud-rest-api/database"
	product_querys "golang-crud-rest-api/querys"
	strconve "golang-crud-rest-api/string_conversion"
	products "golang-crud-rest-api/types"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutesForProducts(router *mux.Router) {

	EnableCORS(router)

	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		product, err := product_querys.GetProduct()
		if err == nil {
			RespondWithSuccess(product, w)

		} else {
			RespondWithError(err, w)

		}
	}).Methods(http.MethodGet)

	//To get by id
	router.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := strconve.StringToInt64(idAsString)

		if err != nil {
			RespondWithError(err, w)
			return
		}
		product, err := product_querys.GetProductById(id)

		if err != nil {
			RespondWithError(err, w)
		} else {
			RespondWithSuccess(product, w)
		}

	}).Methods(http.MethodGet)

	//To Create
	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {

		var product products.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			RespondWithError(err, w)
		} else {
			err := product_querys.CreateProduct(product)
			if err != nil {
				RespondWithError(err, w)
			} else {
				RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	//TO Update
	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {

		var product products.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			RespondWithError(err, w)
		} else {
			err := product_querys.UpdateProduct(product)
			if err != nil {
				RespondWithError(err, w)
			} else {
				RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

	//To DELETE
	router.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := strconve.StringToInt64(idAsString)

		if err != nil {
			RespondWithError(err, w)
			return
		}
		err = product_querys.DeleteProduct(id)

		if err != nil {
			RespondWithError(err, w)
		} else {
			RespondWithSuccess(true, w)
		}

	}).Methods(http.MethodDelete)

}

//To enable CORS
func EnableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", database.AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(MiddlewareCors)
}

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", database.AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

//To respond with error when therer is one
func RespondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

//To respond with success when there is no error
func RespondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
