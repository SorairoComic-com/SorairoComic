package routers

import (
	"sorairocomic/controllers"

	"github.com/gin-gonic/gin"
)

func Comics(router *gin.RouterGroup) {
	router.GET("/", controllers.GetComics)
	router.POST("/", controllers.PostComic)
	router.DELETE("/:comicId", controllers.DeleteComic)
	router.PUT("/:comicId", controllers.PutComic)
}
