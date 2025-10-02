package services

import (
	"fmt"

	"github.com/ian0113/go-gin-mvc/config"
	"github.com/ian0113/go-gin-mvc/infra"
	"github.com/ian0113/go-gin-mvc/repositories"
	"github.com/ian0113/go-gin-mvc/utils"

	"go.uber.org/zap"
)

type AuthService struct {
	logger *zap.Logger
	repo   *repositories.AuthRepository
	cfg    *config.AuthServiceConfig
}

func NewAuthService() *AuthService {
	return &AuthService{
		logger: infra.GetLogger().Named("auth.service"),
		repo:   repositories.NewAuthRepository(),
		cfg:    &infra.GetConfig().AuthService,
	}
}

func (x *AuthService) SetRefreshToken(userID uint) (string, error) {
	expiration := x.cfg.RefreshTokenExpiration
	token, err := utils.GenerateToken(userID, expiration)
	if err != nil {
		return "", err
	}
	err = x.repo.SetRefreshToken(token, true, expiration)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (x *AuthService) DelRefreshToken(token string) error {
	return x.repo.DelRefreshToken(token)
}

func (x *AuthService) SetAccessToken(userID uint) (string, error) {
	expiration := x.cfg.AccessTokenExpiration
	token, err := utils.GenerateToken(userID, expiration)
	if err != nil {
		return "", err
	}
	err = x.repo.SetAccessToken(token, userID, expiration)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (x *AuthService) DelAccessToken(token string) error {
	return x.repo.DelAccessToken(token)
}

func (x *AuthService) RefreshAccessToken(refreshToken string) (string, error) {
	jwtRefreshToken, err := utils.ValidateJWT(refreshToken)
	if err != nil {
		return "", err
	}
	userID, err := utils.ExtractUserID(jwtRefreshToken)
	if err != nil {
		return "", err
	}
	allow, err := x.repo.GetRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}
	if !allow {
		return "", fmt.Errorf("not allowed")
	}
	return x.SetAccessToken(userID)
}

func (x *AuthService) ValidateAccessToken(token string) (uint, error) {
	jwtToken, err := utils.ValidateJWT(token)
	if err != nil {
		return 0, err
	}
	userID, err := utils.ExtractUserID(jwtToken)
	if err != nil {
		return 0, err
	}
	_, err = x.repo.GetAccessToken(token)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
