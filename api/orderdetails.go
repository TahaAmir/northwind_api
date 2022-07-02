package api

import (
	"encoding/json"
	"golang-crud-rest-api/middleware"
	"golang-crud-rest-api/querys"
	types "golang-crud-rest-api/type"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRoutesForOrderDetails(router *mux.Router) {

	middleware.EnableCORS(router)

	router.HandleFunc("/orderdetails", func(w http.ResponseWriter, r *http.Request) {
		order_details, err := querys.GetOrderDetails()
		if err == nil {
			middleware.RespondWithSuccess(order_details, w)

		} else {
			middleware.RespondWithSuccess(err, w)
		}

	}).Methods(http.MethodGet)

	router.HandleFunc("/orderdetails/{id}", func(w http.ResponseWriter, r *http.Request) {

		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)

		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		order_details, err := querys.GetOrderDetailsById(id)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(order_details, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/orderdetails", func(w http.ResponseWriter, r *http.Request) {
		var od types.OrderDetails
		err := json.NewDecoder(r.Body).Decode(&od)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.CreateOrderDetails(od)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/orderdetails", func(w http.ResponseWriter, r *http.Request) {
		var od types.OrderDetails
		err := json.NewDecoder(r.Body).Decode(&od)

		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			err := querys.UpdateOrderDetails(od)
			if err != nil {
				middleware.RespondWithError(err, w)
			} else {
				middleware.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

	router.HandleFunc("/orderdetails/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := StringToInt64(idAsString)
		if err != nil {
			middleware.RespondWithError(err, w)
			return
		}
		err = querys.DeleteOrderDetails(id)
		if err != nil {
			middleware.RespondWithError(err, w)
		} else {
			middleware.RespondWithSuccess(true, w)
		}

	}).Methods(http.MethodDelete)

}
