package router

import (
	"go-finance/endpoints"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/", endpoints.Authenticate)
	router.HandleFunc("/createuser", endpoints.AddUserEndPoint).Methods("POST")
	router.HandleFunc("/validate/{userid}", endpoints.ValidateUserEndPoint).Methods("GET")
	router.HandleFunc("/transactions/{userid}", endpoints.AllTransactionsEndPoint).Methods("GET")
	router.HandleFunc("/topup/{userid}", endpoints.TopUpEndPoint).Methods("POST")
	router.HandleFunc("/balance/{userid}", endpoints.GetBalanceEndPoint).Methods("GET")
	http.ListenAndServe(":2021", router)
}
