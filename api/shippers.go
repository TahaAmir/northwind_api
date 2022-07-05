package api

import (
	"encoding/json"
	"fmt"
	"golang-crud-rest-api/middleware"
	"golang-crud-rest-api/querys"
	types "golang-crud-rest-api/type"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutesForShippers(router *mux.Router) {

	middleware.EnableCORS(router)

	router.HandleFunc("/shippers", func(w http.ResponseWriter, r *http.Request) {
		shipper, err := querys.GetShippers()
		if err == nil {
			middleware.RespondWithJSON(w, http.StatusOK, shipper)

		} else {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())

		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/shippers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Shipper ID")
			return
		}
		shipper, err := querys.GetShippersById(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusNotFound, "Shipper's Id not Found")
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, shipper)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/shippers", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var s types.Shippers
		err := json.NewDecoder(r.Body).Decode(&s)
		fmt.Println(s)
		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		} else {
			err := querys.CreateShippers(s)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				middleware.RespondWithJSON(w, http.StatusOK, true)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/shippers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		err = querys.DeleteShippers(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, true)
		}
	}).Methods(http.MethodDelete)

	router.HandleFunc("/shippers", func(w http.ResponseWriter, r *http.Request) {
		var c types.Shippers
		err := json.NewDecoder(r.Body).Decode(&c)
		fmt.Println(c)
		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		} else {
			err := querys.UpdateShippers(c)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				middleware.RespondWithJSON(w, http.StatusOK, true)
			}
		}
	}).Methods(http.MethodPut)
}
