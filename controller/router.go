package controller

import (
	"github.com/gin-gonic/gin"
	authService "github.com/lukmanlukmin/wallet/service/auth"
	transactionService "github.com/lukmanlukmin/wallet/service/transaction"
	userService "github.com/lukmanlukmin/wallet/service/user"
	authMidleware "github.com/lukmanlukmin/wallet/util/middleware"
)

type RouterLoader struct {
}

func LoadRouter(routers *gin.Engine) {
	router := &RouterLoader{}
	router.AuthRouter(routers)
	router.UserRouter(routers)
	router.TransactionRouter(routers)
}

func (rLoader *RouterLoader) AuthRouter(router *gin.Engine) {
	handler := &AuthController{
		AuthService: authService.AuthServiceHandler(),
	}
	rLoader.AuthRouterDefinition(router, handler)
}

func (rLoader *RouterLoader) AuthRouterDefinition(router *gin.Engine, handler *AuthController) {
	authGroup := router.Group("auth")
	authGroup.POST("login", handler.Login)

	midleware := authMidleware.DefaultMiddleware{}
	authGroup.Use(midleware.JWTAuthMidlewareGuest())
	{
		authGroup.GET("logout", handler.Logout)
	}
}

func (rLoader *RouterLoader) TransactionRouter(router *gin.Engine) {
	handler := &TransactionController{
		TransactionService: transactionService.TransactionServiceHandler(),
	}
	rLoader.TransactionRouterDefinition(router, handler)
}

func (rLoader *RouterLoader) TransactionRouterDefinition(router *gin.Engine, handler *TransactionController) {
	transactionGroup := router.Group("transaction")
	midleware := authMidleware.DefaultMiddleware{}
	transactionGroup.Use(midleware.JWTAuthMidlewareGuest())
	{
		transactionGroup.POST("topup", handler.TopUp)
		transactionGroup.POST("transfer", handler.Transfer)
	}
}

func (rLoader *RouterLoader) UserRouter(router *gin.Engine) {
	handler := &UserController{
		UserService: userService.UserServiceHandler(),
	}
	rLoader.UserRouterDefinition(router, handler)
}

func (rLoader *RouterLoader) UserRouterDefinition(router *gin.Engine, handler *UserController) {
	userGroup := router.Group("user")
	midleware := authMidleware.DefaultMiddleware{}

	userGroup.Use(midleware.JWTAuthMidlewareGuest())
	{
		userGroup.GET("", handler.GetUserMe)
		userGroup.PUT("", handler.CreateUser)
	}

	userGroup.Use(midleware.JWTAuthMidlewareAdmin())
	{
		userGroup.POST("", handler.CreateUser)
		userGroup.GET(":id", handler.GetUserByID)
		userGroup.PUT(":id", handler.GetUserByID)
		userGroup.DELETE(":id", handler.GetUserByID)
	}

}
