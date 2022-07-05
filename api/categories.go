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

func SetupRoutesForCategories(router *mux.Router) {

	middleware.EnableCORS(router)

	router.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		category, err := querys.GetCategory()
		if err == nil {
			middleware.RespondWithJSON(w, http.StatusOK, category)

		} else {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())

		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Category ID")
			return
		}
		category, err := querys.GetCategoryById(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusNotFound, "Category not found")
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, category)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var c types.Catogories
		err := json.NewDecoder(r.Body).Decode(&c)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		} else {
			err := querys.CreateCatogorry(c)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				middleware.RespondWithJSON(w, http.StatusCreated, c)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Category ID")
			return
		}
		err = querys.DeleteCatogory(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, true)
		}
	}).Methods(http.MethodDelete)

	router.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		var c types.Catogories
		err := json.NewDecoder(r.Body).Decode(&c)
		fmt.Println(c)
		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		} else {
			err := querys.UpdateCategory(c)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				middleware.RespondWithJSON(w, http.StatusOK, true)
			}
		}
	}).Methods(http.MethodPut)
}
