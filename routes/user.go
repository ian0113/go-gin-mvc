package routes

import (
	"github.com/ian0113/go-gin-mvc/controllers"
	"github.com/ian0113/go-gin-mvc/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	registerApiRouter(func(rg *gin.RouterGroup) {
		authMiddleware := middlewares.NewAuthMiddleware()
		controller := controllers.NewUserController()
		r := rg.Group("/users")
		r.POST("", controller.Register)
		r.DELETE("/:id", authMiddleware.ValidAuthStatus, controller.Unregister)
	})
}
