package types

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginBody struct {
	UsernameOrEmail string
	Password        string
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type Token struct {
	Id primitive.ObjectID `json:"_id"`
	jwt.StandardClaims
}

type RegisterBody struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password,omitempty"`
}
