package main

import (
	"fmt"
	"sorairocomic/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	port := 3000

	routers.Index(app.Group("/"))

	app.Run(fmt.Sprintf(":%d", port))
	fmt.Println("http://localhost:3000/")
}
