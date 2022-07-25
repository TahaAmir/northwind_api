package api

import (
	"encoding/json"
	"fmt"
	"golang-crud-rest-api/middleware"
	"golang-crud-rest-api/querys"
	types "golang-crud-rest-api/type"
	"io/ioutil"
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

		var shipper types.Shippers
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		json.Unmarshal([]byte(body), &shipper)

		res, err := querys.CreateShippers(shipper)
		if err != nil {

			middleware.RespondWithError(w, http.StatusConflict, err.Error())
			return
		}
		middleware.RespondWithJSON(w, http.StatusCreated, res)
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
			middleware.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
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
				middleware.RespondWithJSON(w, http.StatusOK, c)
			}
		}
	}).Methods(http.MethodPut)
}
