package routers

import (
	"github.com/gin-gonic/gin"
)

func Index(router *gin.RouterGroup) {
	User(router.Group("/"))
}
