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
			middleware.RespondWithSuccess(category, w)

		} else {
			middleware.RespondWithError(err, w)

		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		category, err := querys.GetCategoryById(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(category, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var c types.Catogories
		err := json.NewDecoder(r.Body).Decode(&c)
		fmt.Println(c)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.CreateCatogorry(c)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		err = querys.DeleteCatogory(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)

	router.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
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
