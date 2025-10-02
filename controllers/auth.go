package controllers

import (
	"net/http"

	"github.com/ian0113/go-gin-mvc/infra"
	"github.com/ian0113/go-gin-mvc/services"
	"github.com/ian0113/go-gin-mvc/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	logger      *zap.Logger
	userService *services.UserService
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		logger:      infra.GetLogger().Named("auth.controller"),
		userService: services.NewUserService(),
		authService: services.NewAuthService(),
	}
}

// POST /api/auth/login
func (x *AuthController) Login(c *gin.Context) {
	type LoginRequest struct {
		Account  string `json:"account" binding:"required,min=6"`
		Password string `json:"password" binding:"required,min=6"`
	}
	req := LoginRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := x.userService.ValidateUser(req.Account, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid account or password"})
		return
	}

	refreshToken, err := x.authService.SetRefreshToken(user.ID)
	if err != nil {
		x.authService.DelRefreshToken(refreshToken)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set auth status"})
		return
	}

	accessToken, err := x.authService.SetAccessToken(user.ID)
	if err != nil {
		x.authService.DelRefreshToken(refreshToken)
		x.authService.DelAccessToken(accessToken)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set auth status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"refresh_token": refreshToken,
		"access_token":  accessToken,
	})
}

// POST /api/auth/logout
func (x *AuthController) Logout(c *gin.Context) {
	type LogoutRequest struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	req := LogoutRequest{}

	accessToken, ok := utils.GetGinContextAccessToken(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get access token"})
		return
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err = x.authService.DelRefreshToken(req.RefreshToken)
	if err != nil {
		x.logger.Error("Failed to delete refresh token", zap.Error(err))
	}

	err = x.authService.DelAccessToken(accessToken)
	if err != nil {
		x.logger.Error("Failed to delete access token", zap.Error(err))
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

// POST /api/auth/refresh
func (x *AuthController) Refresh(c *gin.Context) {
	type RefreshRequest struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	req := RefreshRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	accessToken, err := x.authService.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
