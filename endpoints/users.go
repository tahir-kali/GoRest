package endpoints

import (
	"encoding/json"
	"go-finance/entities"
	"go-finance/services"
	"net/http"

	"github.com/gorilla/mux"
)

func User(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authenticate"))
}
func ValidateUserEndPoint(response http.ResponseWriter, request *http.Request) {

	data := services.FindAccount(mux.Vars(request)["userid"])
	//if user doesnt exist then create one and create a wallet for the user
	if !data {
		var person entities.User
		person.Email = mux.Vars(request)["userid"]
		person.First = "First"
		person.Last = "Last"
		services.CreateNewAccount(person)
		services.CreateWallet(mux.Vars(request)["userid"])
	}
	json.NewEncoder(response).Encode(data)
}
func AddUserEndPoint(response http.ResponseWriter, request *http.Request) {
	if request.Body != nil {
		var person entities.User
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&person)
		if err != nil {
			panic(err)
		}
		result := services.CreateNewAccount(person)
		response.Write([]byte(result))
	} else {
		response.Write([]byte("Please enter your email, password, first name and last name in request body"))
	}

}
