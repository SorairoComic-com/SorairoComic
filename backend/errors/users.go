package errors

import (
	"net/http"
	"sorairocomic/types"

	"github.com/gin-gonic/gin"
)

func generateMessageError(message string) types.Error {
	return types.Error{
		Message: message,
	}
}

func GetMessageError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, generateMessageError(message))
	return
}
