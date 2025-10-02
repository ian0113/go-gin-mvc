package services

import (
	"github.com/ian0113/go-gin-mvc/infra"
	"github.com/ian0113/go-gin-mvc/models"
	"github.com/ian0113/go-gin-mvc/repositories"

	"go.uber.org/zap"
)

type OrderService struct {
	logger *zap.Logger
	repo   *repositories.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		logger: infra.GetLogger().Named("order.service"),
		repo:   repositories.NewOrderRepository(),
	}
}

func (x *OrderService) CreateOrder(order *models.Order) error {
	return x.repo.Create(order)
}

func (x *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	return x.repo.FindByID(id)
}

func (x *OrderService) GetOrders(limit int) ([]models.Order, error) {
	return x.repo.FindN(limit)
}

func (x *OrderService) UpdateOrder(order *models.Order) error {
	return x.repo.Update(order)
}

func (x *OrderService) DeleteOrder(id uint) error {
	return x.repo.DeleteByID(id)
}
