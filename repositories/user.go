package repositories

import (
	"github.com/ian0113/go-gin-mvc/infra"
	"github.com/ian0113/go-gin-mvc/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		logger: infra.GetLogger().Named("user.repository"),
		db:     infra.GetDB(),
	}
}

func (x *UserRepository) Create(user *models.User) error {
	return x.db.Create(user).Error
}

func (x *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	return &user, x.db.Where("id = ?", id).Find(&user).Error
}

func (x *UserRepository) DeleteByID(id uint) error {
	return x.db.Where("id = ?", id).Delete(&models.User{}).Error
}

func (x *UserRepository) FindByAccount(account string) (*models.User, error) {
	var user models.User
	return &user, x.db.Where("account = ?", account).Find(&user).Error
}

func (x *UserRepository) DeleteByAccounnt(account string) error {
	return x.db.Where("account = ?", account).Delete(&models.User{}).Error
}
