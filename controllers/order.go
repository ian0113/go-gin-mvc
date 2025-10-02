package controllers

import (
	"net/http"
	"strconv"

	"github.com/ian0113/go-gin-mvc/infra"
	"github.com/ian0113/go-gin-mvc/models"
	"github.com/ian0113/go-gin-mvc/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderController struct {
	logger  *zap.Logger
	service *services.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		logger:  infra.GetLogger().Named("order.controller"),
		service: services.NewOrderService(),
	}
}

// POST /api/orders
func (x *OrderController) CreateOrder(c *gin.Context) {
	var order models.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = x.service.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GET /api/orders?limit=10
func (x *OrderController) ListOrders(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	orders, err := x.service.GetOrders(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GET /api/orders/:id
func (x *OrderController) GetOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	order, err := x.service.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// PUT /api/orders/:id
func (x *OrderController) UpdateOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	var order models.Order
	err = c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.ID = uint(id)

	err = x.service.UpdateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// DELETE /api/orders/:id
func (x *OrderController) DeleteOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	err = x.service.DeleteOrder(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}
