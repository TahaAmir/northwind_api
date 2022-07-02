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

func SetUpRoutesForOrders(router *mux.Router) {

	middleware.EnableCORS(router)

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		order, err := querys.GetOrders()
		if err == nil {
			middleware.RespondWithSuccess(order, w)
		} else {
			middleware.RespondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		orders, err := querys.GetOrdersByID(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(orders, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {

		var o types.Orders
		err := json.NewDecoder(r.Body).Decode(&o)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.CreateOrders(o)
			if err != nil {
				middleware.RespondWithError(err, w)
				fmt.Println(err)

			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)
		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		err = querys.DeleteOrders(id)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		var o types.Orders
		err := json.NewDecoder(r.Body).Decode(&o)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.UpdateOrders(o)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)
}
