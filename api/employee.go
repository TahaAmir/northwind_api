package api

import (
	"encoding/json"
	"golang-crud-rest-api/middleware"
	"golang-crud-rest-api/querys"
	types "golang-crud-rest-api/type"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutesForEmployee(router *mux.Router) {
	middleware.EnableCORS(router)

	router.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		employee, err := querys.GetEmployee()
		if err == nil {
			middleware.RespondWithSuccess(employee, w)
		} else {
			middleware.RespondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/employee/{id}", func(w http.ResponseWriter, r *http.Request) {

		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		employee, err := querys.GetEmployeeByID(id)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(employee, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		var e types.Employee
		err := json.NewDecoder(r.Body).Decode(&e)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.CreateEmployee(e)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		var e types.Employee
		err := json.NewDecoder(r.Body).Decode(&e)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.UpdateEmployee(e)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

	router.HandleFunc("/employee/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)
		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		err = querys.DeleteEmployee(id)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(true, w)
		}

	}).Methods(http.MethodDelete)
}
