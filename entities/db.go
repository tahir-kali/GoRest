package entities

import "cloud.google.com/go/firestore"

type Database struct {
	Client *firestore.Client
}
