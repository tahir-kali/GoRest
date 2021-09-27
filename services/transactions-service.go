package services

import (
	"context"
	"go-finance/entities"
	"log"

	"google.golang.org/api/iterator"
)

func GetAllTransactions(email string) []entities.Transaction {
	var transaction []entities.Transaction
	dsnap := FireApp().Collection("transactions").Where("User", "==", email).Documents(context.Background())
	for {
		doc, err := dsnap.Next()
		if err == iterator.Done {
			break
		}
		// manage, _ := json.Marshal(doc.Data())
		var operation entities.Transaction
		// json.Unmarshal(manage, res)
		doc.DataTo(&operation)
		transaction = append(transaction, operation)
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

	}
	return transaction
}
func RecordTransaction(transaction entities.Transaction) {

	FireApp().Collection("transactions").Add(context.Background(), transaction)
}
