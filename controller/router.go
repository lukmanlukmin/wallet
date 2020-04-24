package controller

import (
	"github.com/gin-gonic/gin"
	authService "github.com/lukmanlukmin/wallet/service/auth"
	userService "github.com/lukmanlukmin/wallet/service/user"
	authMidleware "github.com/lukmanlukmin/wallet/util/middleware"
)

type RouterLoader struct {
}

func LoadRouter(routers *gin.Engine) {
	router := &RouterLoader{}
	router.AuthRouter(routers)
	router.UserRouter(routers)
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
