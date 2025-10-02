package routes

import (
	"github.com/ian0113/go-gin-mvc/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	r.Use(cors.Default())

	apiMiddleware := middlewares.NewApiMiddleware()
	api := r.Group("/api")
	api.Use(apiMiddleware.Logger)
	runRegisterApiRouters(api)
}
