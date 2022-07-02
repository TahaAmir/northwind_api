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
			middleware.RespondWithSuccess(shipper, w)

		} else {
			middleware.RespondWithError(err, w)

		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/shippers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		shipper, err := querys.GetShippersById(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(shipper, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/shippers", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var s types.Shippers
		err := json.NewDecoder(r.Body).Decode(&s)
		fmt.Println(s)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.CreateShippers(s)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/shippers/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		err = querys.DeleteShippers(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)

	router.HandleFunc("/shippers", func(w http.ResponseWriter, r *http.Request) {
		var c types.Catogories
		err := json.NewDecoder(r.Body).Decode(&c)
		fmt.Println(c)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.UpdateCategory(c)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)
}
