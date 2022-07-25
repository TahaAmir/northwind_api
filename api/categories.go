package api

import (
	"encoding/json"
	"golang-crud-rest-api/middleware"
	"golang-crud-rest-api/querys"
	types "golang-crud-rest-api/type"
	"io/ioutil"
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

	//
	//
	//

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

	//
	//
	//

	router.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		var category types.Catogories
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		json.Unmarshal([]byte(body), &category)

		res, err := querys.CreateCatogorry(category)
		if err != nil {

			middleware.RespondWithError(w, http.StatusConflict, err.Error())
			return
		}
		middleware.RespondWithJSON(w, http.StatusCreated, res)
	}).Methods(http.MethodPost)
	//
	//
	//
	router.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {

		var c types.Catogories
		err := json.NewDecoder(r.Body).Decode(&c)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		} else {
			err := querys.UpdateCategory(c)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				middleware.RespondWithJSON(w, http.StatusOK, c)
			}
		}
	}).Methods(http.MethodPut)

	//
	//
	//
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
			middleware.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
		}
	}).Methods(http.MethodDelete)

}
