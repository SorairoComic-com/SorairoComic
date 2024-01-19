package errors

import (
	"sorairocomic/types"

	"github.com/gin-gonic/gin"
)

func generateMessageError(message string) types.Error {
	return types.Error{
		Message: message,
	}
}

func GetMessageError(c *gin.Context, status int, message string) {
	c.JSON(status, generateMessageError(message))
	return
}
