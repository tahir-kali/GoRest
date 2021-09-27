package services

import (
	"context"
	"go-finance/entities"
	"log"

	"cloud.google.com/go/firestore"
)

func TopUpWallet(transaction entities.Transaction) {
	var currentBalance = GetBalance(transaction.User)
	FireApp().Collection("wallets").Doc(transaction.User).Update(context.Background(), []firestore.Update{
		{
			Path:  "balance",
			Value: currentBalance + transaction.Balance,
		},
	})
	RecordTransaction(transaction)
}
func GetBalance(email string) int {
	var wallet entities.Wallet
	dsnap, err := FireApp().Collection("wallets").Doc(email).Get(context.Background())

	if err != nil {
		log.Fatal(err.Error())
	}
	if dsnap.Exists() {
		dsnap.DataTo(&wallet)
	}
	return wallet.Balance
}
func CreateWallet(email string) {
	var wallet entities.Wallet
	wallet.Email = email
	wallet.Balance = 0
	FireApp().Collection("wallets").Doc(email).Set(context.Background(), wallet)
}
