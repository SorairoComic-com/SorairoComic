package controllers

import (
	"context"
	"net/http"

	"sorairocomic/config"
	"sorairocomic/errors"
	"sorairocomic/helpers"
	"sorairocomic/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var coll = config.MongoDB().Collection("Users")

func Login(c *gin.Context) {
	var requestBody types.LoginBody
	if err := c.BindJSON(&requestBody); err != nil {
		errors.GetMessageError(c, "Invalid Body")
		return
	}

	var users []types.User
	pipeline := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "$or", Value: []bson.D{
				{{Key: "username", Value: requestBody.UsernameOrEmail}},
				{{Key: "email", Value: requestBody.UsernameOrEmail}},
			}},
		}},
	}

	cursor, _ := coll.Aggregate(context.TODO(), mongo.Pipeline{pipeline})
	if err := cursor.All(context.TODO(), &users); err != nil {
		errors.GetMessageError(c, "Invalid Username/Email")
		return
	}

	if len(users) == 0 {
		errors.GetMessageError(c, "Invalid Username/Email")
		return
	}

	if isLogin := helpers.ComparePassword(requestBody.Password, users[0].Password); isLogin == false {
		errors.GetMessageError(c, "Invalid Username/Email")
		return
	}

	token := helpers.CreateToken(types.Token{
		Id: users[0].Id,
	})

	c.JSON(http.StatusOK, types.LoginResponse{
		AccessToken: token,
	})
}

func Register(c *gin.Context) {
	var requestBody types.RegisterBody
	if err := c.BindJSON(&requestBody); err != nil {
		errors.GetMessageError(c, "Invalid Body")
		return
	}

	v := validator.New()
	if err := v.Struct(requestBody); err != nil {
		if err.(validator.ValidationErrors)[0].Tag() == "required" {
			errors.GetMessageError(c, err.(validator.ValidationErrors)[0].Field()+" is required")
			return
		}
		if err.(validator.ValidationErrors)[0].Tag() == "email" {
			errors.GetMessageError(c, "Invalid email format")
			return
		}
		if err.(validator.ValidationErrors)[0].Tag() == "min" {
			errors.GetMessageError(c, "Password should at least 8 characters")
			return
		}
	}

	var users []types.User
	pipeline := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "$or", Value: []bson.D{
				{{Key: "username", Value: requestBody.Username}},
				{{Key: "email", Value: requestBody.Email}},
			}},
		}},
	}

	cursor, _ := coll.Aggregate(context.TODO(), mongo.Pipeline{pipeline})
	if err := cursor.All(context.TODO(), &users); err == nil && len(users) > 0 {
		if users[0].Username == requestBody.Username {
			errors.GetMessageError(c, "Username already exists")
			return
		}
		if users[0].Email == requestBody.Email {
			errors.GetMessageError(c, "Email already exists")
			return
		}
	}

	hashedPassword := helpers.HashPassword(requestBody.Password)

	result, _ := coll.InsertOne(context.TODO(), types.RegisterBody{
		Username: requestBody.Username,
		Email:    requestBody.Email,
		Password: hashedPassword,
	})

	c.JSON(http.StatusOK, types.User{
		Id:       result.InsertedID.(primitive.ObjectID),
		Username: requestBody.Username,
		Email:    requestBody.Email,
	})
}
