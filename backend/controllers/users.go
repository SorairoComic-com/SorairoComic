package controllers

import (
	"net/http"
	"sorairocomic/types"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, types.Login{
		AccessToken: "Oke Login",
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, "register")
}
