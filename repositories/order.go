package repositories

import (
	"github.com/ian0113/go-gin-mvc/infra"
	"github.com/ian0113/go-gin-mvc/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		logger: infra.GetLogger().Named("order.repository"),
		db:     infra.GetDB(),
	}
}

func (x *OrderRepository) Create(order *models.Order) error {
	return x.db.Create(order).Error
}

func (x *OrderRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order
	return &order, x.db.Where("id = ?").Find(&order).Error
}

func (x *OrderRepository) FindN(limit int) ([]models.Order, error) {
	var orders []models.Order
	return orders, x.db.Limit(limit).Order("created_at desc").Find(&orders).Error
}

func (x *OrderRepository) Update(order *models.Order) error {
	return x.db.Save(order).Error
}

func (x *OrderRepository) DeleteByID(id uint) error {
	return x.db.Where("id = ?", id).Delete(&models.Order{}).Error
}
