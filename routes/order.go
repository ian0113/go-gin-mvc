package routes

import (
	"github.com/ian0113/go-gin-mvc/controllers"
	"github.com/ian0113/go-gin-mvc/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	registerApiRouter(func(rg *gin.RouterGroup) {
		authMiddleware := middlewares.NewAuthMiddleware()
		controller := controllers.NewOrderController()
		r := rg.Group("/orders")
		r.POST("", authMiddleware.ValidAuthStatus, controller.CreateOrder)
		r.GET("", authMiddleware.ValidAuthStatus, controller.ListOrders)
		r.GET("/:id", authMiddleware.ValidAuthStatus, controller.GetOrder)
		r.PUT("/:id", authMiddleware.ValidAuthStatus, controller.UpdateOrder)
		r.DELETE("/:id", authMiddleware.ValidAuthStatus, controller.DeleteOrder)
	})
}
