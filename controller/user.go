package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	httpRequestEntity "github.com/lukmanlukmin/wallet/entity/http/request"
	service "github.com/lukmanlukmin/wallet/service/user"
)

type UserController struct {
	UserService service.UserServiceInterface
}

func (handler *UserController) GetUserMe(context *gin.Context) {
	walletId := context.Request.Header["Wallet-Uid"][0]
	idUser, _ := strconv.Atoi(walletId)
	result, _ := handler.UserService.GetUserByID(idUser)
	context.JSON(http.StatusOK, result)
}

func (handler *UserController) GetUserByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	result, err := handler.UserService.GetUserByID(id)
	if result == nil {
		context.JSON(http.StatusOK, gin.H{})
		return
	}
	context.JSON(http.StatusOK, result)
}

func (handler *UserController) CreateUser(context *gin.Context) {

	bodyPayload := httpRequestEntity.UserRequest{}
	context.ShouldBind(&bodyPayload)

	err := handler.UserService.CreateUser(bodyPayload)
	if err != nil {
		context.JSON(http.StatusCreated, bodyPayload)
	}
}

func (service *UserController) GetUsers(context *gin.Context) {
	queryparam := &httpRequestEntity.Limitofset{}
	err := context.ShouldBindQuery(queryparam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	result := service.UserService.GetUserList(queryparam.Limit, queryparam.Offset)
	context.JSON(http.StatusOK, result)
}

func (handler *UserController) DeleteUserByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	err2 := handler.UserService.DeleteUserByID(id)
	if err2 == nil {
		context.JSON(http.StatusAccepted, gin.H{})
		return
	}
	context.JSON(http.StatusInternalServerError, "")
}

func (handler *UserController) UpdateUserByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	bodyPayload := httpRequestEntity.UpdateUserRequest{}
	context.ShouldBind(&bodyPayload)

	err2 := handler.UserService.UpdateUser(id, bodyPayload)
	if err2 == nil {
		context.JSON(http.StatusAccepted, gin.H{})
		return
	}
	context.JSON(http.StatusInternalServerError, "")
}
