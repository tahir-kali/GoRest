package services

import (
	"context"
	"go-finance/entities"
)

func CreateNewAccount(data entities.User) string {

	if !FindAccount(data.Email) {
		FireApp().Collection("users").Doc(data.Email).Set(context.Background(), data)
		CreateWallet(data.Email)
		return "User Added"
	} else {
		return "User Exisits"
	}
}
func FindAccount(email string) bool {

	dsnap, err := FireApp().Collection("users").Doc(email).Get(context.Background())
	if err != nil {
		return false
	}
	if dsnap.Exists() {
		return true
	}
	return false
}
