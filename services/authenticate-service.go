package services

import (
	"go-finance/entities"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

func CreateToken(userId string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", string(jwtKey)) //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(token string) bool {
	claims := &entities.Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			//	w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		//w.WriteHeader(http.StatusBadRequest)
		return false
	}

	if !tkn.Valid {
		//w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}
