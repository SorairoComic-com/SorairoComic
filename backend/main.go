package main

import (
	"fmt"
	"sorairocomic/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	port := 3000

	router.PostLogin(app.Group("/login"))

	app.Run(fmt.Sprintf(":%d", port))
	fmt.Println("http://localhost:3000/")
}
