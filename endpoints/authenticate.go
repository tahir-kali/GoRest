package endpoints

import (
	"context"
	"go-finance/entities"
	"go-finance/services"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	userId := r.Header.Get("UserID")
	password := r.Header.Get("Password")
	dsnap, _ := services.FireApp().Collection("users").Doc(userId).Get(context.Background())
	if dsnap.Exists() {
		dsnap.DataTo(&user)
		if user.Password == password {
			token, _ := services.CreateToken(userId)
			w.Header().Add("X-Digest", token)
			w.Write([]byte("Ok"))
		} else {
			w.Write([]byte("Incorrect Password"))
		}
	} else {
		w.Write([]byte("UserID is not found in the database"))
	}

}
