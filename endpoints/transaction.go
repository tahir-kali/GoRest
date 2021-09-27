package endpoints

import (
	"encoding/json"
	"go-finance/services"
	"net/http"

	"github.com/gorilla/mux"
)

func AllTransactionsEndPoint(response http.ResponseWriter, request *http.Request) {
	token := services.ValidateToken(request.Header.Get("X-Digest"))
	if token {
		data := services.GetAllTransactions(mux.Vars(request)["userid"])
		json.NewEncoder(response).Encode(data)
	} else {
		response.Write([]byte("Go Away! You are not authenticated!"))
	}

}
