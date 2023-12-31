package types

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginBody struct {
	Username string
	Password string
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string
	Password string
}

type Token struct {
	Id primitive.ObjectID `json:"_id"`
	jwt.StandardClaims
}

type Error struct {
	Message string `json:"message"`
}
