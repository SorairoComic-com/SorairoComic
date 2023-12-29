package helpers

import (
	"os"
	"sorairocomic/types"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var err = godotenv.Load(".env")

func CreateToken(input types.Token) string {
	if err != nil {
		return "Error ENV"
	}

	var token *jwt.Token

	secret := []byte(os.Getenv("JWT_SECRET"))
	token = jwt.NewWithClaims(jwt.SigningMethodHS256,
		types.Token{
			Id:   input.Id,
			Name: input.Name,
			StandardClaims: jwt.StandardClaims{
				IssuedAt: time.Now().Unix(),
			},
		})

	access_token, _ := token.SignedString(secret)

	return access_token
}
