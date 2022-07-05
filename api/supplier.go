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

func SetupRoutesForSupplier(router *mux.Router) {
	middleware.EnableCORS(router)

	router.HandleFunc("/suppliers", func(w http.ResponseWriter, r *http.Request) {
		count, _ := strconv.Atoi(r.FormValue("count"))
		start, _ := strconv.Atoi(r.FormValue("start"))

		if count > 10 || count < 1 {
			count = 10
		}
		if start < 0 {
			start = 0
		}
		supplier, err := querys.GetSupplier(start, count)

		if err == nil {
			middleware.RespondWithJSON(w, http.StatusOK, supplier)

		} else {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}

	}).Methods(http.MethodGet)

	router.HandleFunc("/suppliers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Supplier ID")
			return
		}
		supplier, err := querys.GetSupplierByID(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusNotFound, "Supplier Not Found")
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, supplier)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/suppliers", func(w http.ResponseWriter, r *http.Request) {

		var s types.Supplier
		err := json.NewDecoder(r.Body).Decode(&s)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		} else {
			err := querys.CreateSuppliers(s)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())

			} else {
				middleware.RespondWithJSON(w, http.StatusOK, true)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/suppliers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)
		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Supplier ID")
			return
		}
		err = querys.DeleteSupplier(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, true)
		}
	}).Methods(http.MethodDelete)

	router.HandleFunc("/suppliers", func(w http.ResponseWriter, r *http.Request) {
		var s types.Supplier
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		} else {
			err := querys.UpdateSuppliers(s)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				middleware.RespondWithJSON(w, http.StatusOK, true)
			}
		}
	}).Methods(http.MethodPut)

}
