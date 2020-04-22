package config

import (
	"github.com/gin-gonic/gin"
	"github.com/lukmanlukmin/wallet/util/middleware"
)

func SetupRouter() *gin.Engine {
	defaultMiddleware := middleware.DefaultMiddleware{}
	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())
	return router
}
