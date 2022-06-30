package api

import (
	"encoding/json"
	"golang-crud-rest-api/middleware"
	"golang-crud-rest-api/querys"
	types "golang-crud-rest-api/type"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SetupRoutesForCustomers(router *mux.Router) {
	middleware.EnableCORS(router)

	router.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		customer, err := querys.GetCustomer()
		if err == nil {
			middleware.RespondWithSuccess(customer, w)
		} else {
			middleware.RespondWithError(err, w)
		}

	}).Methods(http.MethodGet)

	router.HandleFunc("/customers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		customer, err := querys.GetCustomerById(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(customer, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {

		var c types.Customers
		err := json.NewDecoder(r.Body).Decode(&c)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.CreateCustomer(c)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/customers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)
		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		err = querys.DeleteCustomer(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)

	router.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		var c types.Customers
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.UpdateCustomer(c)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

}

func StringToInt64(s string) (int64, error) {
	num, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return num, err
}
