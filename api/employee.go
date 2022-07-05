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

func SetupRoutesForEmployee(router *mux.Router) {
	middleware.EnableCORS(router)

	router.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {

		count, _ := strconv.Atoi(r.FormValue("count"))
		start, _ := strconv.Atoi(r.FormValue("start"))

		if count > 10 || count < 1 {
			count = 10
		}
		if start < 0 {
			start = 0
		}
		employee, err := querys.GetEmployee(start, count)
		if err == nil {
			middleware.RespondWithJSON(w, http.StatusOK, employee)
		} else {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/employee/{id}", func(w http.ResponseWriter, r *http.Request) {

		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Employee ID")
			return
		}
		employee, err := querys.GetEmployeeByID(id)
		if err != nil {
			middleware.RespondWithError(w, http.StatusNotFound, "Employee not found")
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, employee)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		var e types.Employee
		err := json.NewDecoder(r.Body).Decode(&e)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		} else {
			err := querys.CreateEmployee(e)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				middleware.RespondWithJSON(w, http.StatusOK, e)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		var e types.Employee
		err := json.NewDecoder(r.Body).Decode(&e)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		} else {
			err := querys.UpdateEmployee(e)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				middleware.RespondWithJSON(w, http.StatusOK, true)
			}
		}
	}).Methods(http.MethodPut)

	router.HandleFunc("/employee/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)
		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Category ID")
			return
		}
		err = querys.DeleteEmployee(id)
		if err != nil {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, true)
		}

	}).Methods(http.MethodDelete)
}
