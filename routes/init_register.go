package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	globalApiRouters []func(rg *gin.RouterGroup)
)

func registerApiRouter(f func(rg *gin.RouterGroup)) {
	globalApiRouters = append(globalApiRouters, f)
}

func runRegisterApiRouters(rg *gin.RouterGroup) {
	for _, f := range globalApiRouters {
		f(rg)
	}
}
