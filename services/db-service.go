package services

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"

	//"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func FireApp() *firestore.Client {

	var client *firestore.Client
	if client == nil {
		ctx := context.Background()
		sa := option.WithCredentialsFile("config/serviceaccount.json")
		app, err := firebase.NewApp(ctx, nil, sa)
		if err != nil {
			log.Fatalln(err)
		}

		client, err = app.Firestore(ctx)
		if err != nil {
			fmt.Println("Error")
		}

		if err != nil {
			log.Fatalln(err)
		}
		//defer client.Close()
		return client
	} else {
		return client
	}
}
