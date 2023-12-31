package controllers

import (
	"context"
	"net/http"

	"sorairocomic/config"
	"sorairocomic/helpers"
	"sorairocomic/types"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var coll = config.MongoDB().Collection("Users")

func Login(c *gin.Context) {
	var requestBody types.LoginBody
	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
	}

	var user types.User
	err := coll.FindOne(context.TODO(), bson.D{{
		Key: "username", Value: requestBody.Username,
	}}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, types.Error{
			Message: "Invalid Username/Password",
		})
		return
	}

	if user.Password != requestBody.Password {
		c.JSON(http.StatusBadRequest, types.Error{
			Message: "Invalid Username/Password",
		})
		return
	}

	token := helpers.CreateToken(types.Token{
		Id: user.Id,
	})

	c.JSON(http.StatusOK, types.LoginResponse{
		AccessToken: token,
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, "register")
}
