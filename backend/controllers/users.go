package controllers

import (
	"net/http"
	"sorairocomic/helpers"
	"sorairocomic/types"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	token := helpers.CreateToken(types.Token{
		Id:   1,
		Name: "Pandu",
	})

	c.JSON(http.StatusOK, types.Login{
		AccessToken: token,
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, "register")
}
