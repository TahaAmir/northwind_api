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

		count, _ := strconv.Atoi(r.FormValue("count"))
		start, _ := strconv.Atoi(r.FormValue("start"))

		if count > 10 || count < 1 {
			count = 10
		}
		if start < 0 {
			start = 0
		}

		customer, err := querys.GetCustomer(start, count)
		if err == nil {
			middleware.RespondWithJSON(w, http.StatusOK, customer)
		} else {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}

	}).Methods(http.MethodGet)

	router.HandleFunc("/customers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
			return
		}
		customer, err := querys.GetCustomerById(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusNotFound, "Customer not found")
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, customer)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {

		var c types.Customers
		err := json.NewDecoder(r.Body).Decode(&c)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		} else {
			err := querys.CreateCustomer(c)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())

			} else {
				middleware.RespondWithJSON(w, http.StatusOK, true)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/customers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)
		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
			return
		}
		err = querys.DeleteCustomer(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, true)
		}
	}).Methods(http.MethodDelete)

	router.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		var c types.Customers
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		} else {
			err := querys.UpdateCustomer(c)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				middleware.RespondWithJSON(w, http.StatusOK, true)
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
