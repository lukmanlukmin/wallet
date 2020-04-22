package controller

import (
	"github.com/gin-gonic/gin"
)

type RouterLoader struct {
}

func LoadRouter(routers *gin.Engine) {
	router := &RouterLoader{}
	router.AuthRouter(routers)
}

func (rLoader *RouterLoader) AuthRouter(router *gin.Engine) {
	handler := &AuthController{
		// 	UserService: srv.UserServiceHandler(),
	}
	rLoader.routerDefinition(router, handler)
}

func (rLoader *RouterLoader) routerDefinition(router *gin.Engine, handler *AuthController) {
	group := router.Group("auth")
	group.GET("", handler.TestFunction)
}
