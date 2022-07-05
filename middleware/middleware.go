package middleware

import (
	"encoding/json"
	"golang-crud-rest-api/database"
	"net/http"

	"github.com/gorilla/mux"
)

func EnableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", database.AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(MiddlewareCors)
}

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", database.AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

//To respond with error when therer is one
func RespondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	//json.NewEncoder(w).Encode(err.Error())
	RespondWithJSON(w, code, map[string]string{"error": message})
}

//To respond with success when there is no error
func RespondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
