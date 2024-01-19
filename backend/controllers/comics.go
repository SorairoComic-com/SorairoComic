package controllers

import (
	"context"
	"net/http"
	"sorairocomic/config"
	"sorairocomic/errors"
	"sorairocomic/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collComics = config.MongoDB().Collection("Comics")

func GetComics(c *gin.Context) {
	var comics []types.Comics
	cursor, _ := collComics.Find(context.TODO(), bson.D{})
	if err := cursor.All(context.TODO(), &comics); err != nil {
		errors.GetMessageError(c, 500, "Internal Server Error")
		return
	}
	if comics == nil {
		comics = make([]types.Comics, 0)
	}
	c.JSON(http.StatusOK, comics)
}

func PostComic(c *gin.Context) {
	var requestBody types.PostComicBody
	if err := c.BindJSON(&requestBody); err != nil {
		errors.GetMessageError(c, 400, "Invalid Body")
		return
	}

	v := validator.New()
	if err := v.Struct(requestBody); err != nil {
		if err.(validator.ValidationErrors)[0].Tag() == "required" {
			errors.GetMessageError(c, 400, err.(validator.ValidationErrors)[0].Field()+" is required")
			return
		}
	}

	result, _ := collComics.InsertOne(context.TODO(), requestBody)

	c.JSON(http.StatusOK, types.Comics{
		Id:          result.InsertedID.(primitive.ObjectID),
		Name:        requestBody.Name,
		ImageUrl:    requestBody.ImageUrl,
		Price:       requestBody.Price,
		Description: requestBody.Description,
		Genre:       requestBody.Genre,
	})
}

func PutComic(c *gin.Context) {
	comicId := c.Param("comicId")
	c.JSON(http.StatusOK, map[string]string{
		"message": "masuk PUT",
		"comicId": comicId,
	})
}

func DeleteComic(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "masuk DELETE",
	})
}
