package types

import "github.com/golang-jwt/jwt"

type Login struct {
	AccessToken string `json:"access_token"`
}

type Token struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}
