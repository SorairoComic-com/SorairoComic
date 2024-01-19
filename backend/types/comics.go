package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comics struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name"`
	ImageUrl    string             `json:"imgUrl"`
	Price       int                `json:"price"`
	Description string             `json:"description"`
	Genre       string             `json:"genre"`
}

type PostComicBody struct {
	Name        string `validate:"required"`
	ImageUrl    string `validate:"required"`
	Price       int    `validate:"required"`
	Description string `validate:"required"`
	Genre       string `validate:"required"`
}
