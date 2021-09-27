package endpoints

import (
	"encoding/json"
	"go-finance/entities"
	"go-finance/services"
	"net/http"

	"github.com/gorilla/mux"
)

func Wallets(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authenticate"))
}
func TopUpEndPoint(response http.ResponseWriter, request *http.Request) {
	token := services.ValidateToken(request.Header.Get("X-Digest"))
	if token {
		var transaction entities.Transaction
		json.NewDecoder(request.Body).Decode(&transaction)

		services.TopUpWallet(transaction)
	} else {
		response.Write([]byte("Go Away! You are not authenticated!"))
	}
}
func GetBalanceEndPoint(response http.ResponseWriter, request *http.Request) {
	token := services.ValidateToken(request.Header.Get("X-Digest"))
	if token {
		json.NewEncoder(response).Encode(services.GetBalance(mux.Vars(request)["userid"]))
	} else {
		response.Write([]byte("Go Away! You are not authenticated!"))
	}
}
