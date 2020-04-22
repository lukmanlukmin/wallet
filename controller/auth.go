package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func (handler *AuthController) TestFunction(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
