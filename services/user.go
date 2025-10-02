package services

import (
	"fmt"

	"github.com/ian0113/go-gin-mvc/infra"
	"github.com/ian0113/go-gin-mvc/models"
	"github.com/ian0113/go-gin-mvc/repositories"
	"github.com/ian0113/go-gin-mvc/utils"

	"go.uber.org/zap"
)

type UserService struct {
	logger *zap.Logger
	repo   *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		logger: infra.GetLogger().Named("user.service"),
		repo:   repositories.NewUserRepository(),
	}
}

func (x *UserService) CreateUser(name, email, account, password string) (*models.User, error) {
	hash, err := utils.GeneratePasswordHash(password)
	if err != nil {
		return nil, err
	}
	user := models.User{
		Name:     name,
		Email:    email,
		Account:  account,
		Password: hash,
	}
	return &user, x.repo.Create(&user)
}

func (x *UserService) ValidateUser(account, password string) (*models.User, error) {
	user, err := x.repo.FindByAccount(account)
	if err != nil {
		return user, err
	}
	if !utils.CheckPasswordHash(user.Password, password) {
		return user, fmt.Errorf("incorrect password")
	}
	return user, nil
}

func (x *UserService) DeleteUser(id uint) error {
	return x.repo.DeleteByID(id)
}
