package helpers

import (
	"sorairocomic/environments"
	"sorairocomic/types"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(input types.Token) string {
	var token *jwt.Token

	secret := []byte(environments.Environments("JWT_SECRET"))
	token = jwt.NewWithClaims(jwt.SigningMethodHS256,
		types.Token{
			Id: input.Id,
			StandardClaims: jwt.StandardClaims{
				IssuedAt: time.Now().Unix(),
			},
		})

	access_token, _ := token.SignedString(secret)

	return access_token
}
