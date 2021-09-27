package entities

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	UserID   string `json:"email"`
	Password string `json:"password"`
}
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
