package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	httpRequestEntity "github.com/lukmanlukmin/wallet/entity/http/request"
	service "github.com/lukmanlukmin/wallet/service/transaction"
)

type TransactionController struct {
	TransactionService service.TransactionServiceInterface
}

func (handler *TransactionController) TopUp(context *gin.Context) {
	bodyPayload := httpRequestEntity.TopUpRequest{}
	context.ShouldBind(&bodyPayload)
	ipAddr := context.ClientIP()
	userAgent := context.Request.Header.Get("User-Agent")
	agent, _ := json.Marshal(userAgent)
	err := handler.TransactionService.TopUp(1, ipAddr, string(agent), bodyPayload)
	if err != nil {
		context.JSON(http.StatusCreated, "")
	} else {
		context.JSON(http.StatusInternalServerError, "")
	}
}

func (handler *TransactionController) Transfer(context *gin.Context) {
	bodyPayload := httpRequestEntity.TransferRequest{}
	context.ShouldBind(&bodyPayload)
	ipAddr := context.ClientIP()
	userAgent := context.Request.Header.Get("User-Agent")
	agent, _ := json.Marshal(userAgent)
	err := handler.TransactionService.Transfer(1, ipAddr, string(agent), bodyPayload)
	if err != nil {
		context.JSON(http.StatusCreated, "")
	} else {
		context.JSON(http.StatusInternalServerError, "")
	}
}
