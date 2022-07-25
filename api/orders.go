package api

import (
	"encoding/json"
	"golang-crud-rest-api/middleware"
	"golang-crud-rest-api/querys"
	types "golang-crud-rest-api/type"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SetUpRoutesForOrders(router *mux.Router) {

	middleware.EnableCORS(router)

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {

		count, _ := strconv.Atoi(r.FormValue("count"))
		start, _ := strconv.Atoi(r.FormValue("start"))

		if count >= 10 || count <= 1 {
			count = 10
		}
		if start < 0 {
			start = 0
		}
		order, err := querys.GetOrders(start, count)
		if err == nil {
			middleware.RespondWithJSON(w, http.StatusOK, order)
		} else {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Order ID")
			return
		}
		orders, err := querys.GetOrdersByID(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusNotFound, "Order Not found")
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, orders)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {

		var order types.Orders
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		json.Unmarshal([]byte(body), &order)

		res, err := querys.CreateOrders(order)
		if err != nil {

			middleware.RespondWithError(w, http.StatusConflict, err.Error())
			return
		}
		middleware.RespondWithJSON(w, http.StatusCreated, res)
	}).Methods(http.MethodPost)

	router.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)
		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid Order ID")
			return
		}
		err = querys.DeleteOrders(id)

		if err != nil {
			middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			middleware.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
		}
	}).Methods(http.MethodDelete)

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		var o types.Orders
		err := json.NewDecoder(r.Body).Decode(&o)
		if err != nil {
			middleware.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		} else {
			err := querys.UpdateOrders(o)
			if err != nil {
				middleware.RespondWithError(w, http.StatusInternalServerError, err.Error())
			} else {
				middleware.RespondWithJSON(w, http.StatusOK, o)
			}
		}
	}).Methods(http.MethodPut)
}
