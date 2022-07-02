package api

import (
	"encoding/json"
	"golang-crud-rest-api/middleware"
	"golang-crud-rest-api/querys"
	types "golang-crud-rest-api/type"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutesForSupplier(router *mux.Router) {
	middleware.EnableCORS(router)

	router.HandleFunc("/suppliers", func(w http.ResponseWriter, r *http.Request) {
		supplier, err := querys.GetSupplier()
		if err == nil {
			middleware.RespondWithSuccess(supplier, w)
		} else {
			middleware.RespondWithError(err, w)
		}

	}).Methods(http.MethodGet)

	router.HandleFunc("/suppliers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		supplier, err := querys.GetSupplierByID(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(supplier, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/suppliers", func(w http.ResponseWriter, r *http.Request) {

		var s types.Supplier
		err := json.NewDecoder(r.Body).Decode(&s)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.CreateSuppliers(s)
			if err != nil {
				middleware.RespondWithError(err, w)

			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/suppliers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)
		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		err = querys.DeleteSupplier(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)

	router.HandleFunc("/suppliers", func(w http.ResponseWriter, r *http.Request) {
		var s types.Supplier
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.UpdateSuppliers(s)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

}
