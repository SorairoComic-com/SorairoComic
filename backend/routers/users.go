package routers

import (
	"sorairocomic/controllers"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)
}
