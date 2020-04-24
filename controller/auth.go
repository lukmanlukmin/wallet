package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	httpRequestEntity "github.com/lukmanlukmin/wallet/entity/http/request"
	service "github.com/lukmanlukmin/wallet/service/auth"
)

type AuthController struct {
	AuthService service.AuthServiceInterface
}

func (handler *AuthController) TestFunction(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func (handler *AuthController) Login(context *gin.Context) {
	bodyPayload := httpRequestEntity.LoginRequest{}
	context.ShouldBind(&bodyPayload)
	response, err := handler.AuthService.Login(bodyPayload.Email, bodyPayload.Password)
	if err == nil {
		context.JSON(http.StatusCreated, response)
	} else {
		context.JSON(http.StatusUnauthorized, "")
	}
}

func (handler *AuthController) Logout(context *gin.Context) {
	walletId := context.Request.Header["wallet-uid"][0]
	idUser, err := strconv.Atoi(walletId)
	go handler.AuthService.Logout(idUser)
	if err == nil {
		context.JSON(http.StatusCreated, "")
	} else {
		context.JSON(http.StatusUnauthorized, "")
	}
}
