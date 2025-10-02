package middlewares

import (
	"net/http"
	"strings"

	"github.com/ian0113/go-gin-mvc/infra"
	"github.com/ian0113/go-gin-mvc/services"
	"github.com/ian0113/go-gin-mvc/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthMiddleware struct {
	logger      *zap.Logger
	authService *services.AuthService
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		logger:      infra.GetLogger().Named("auth.middleware"),
		authService: services.NewAuthService(),
	}
}

func (x *AuthMiddleware) ValidAuthStatus(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
		return
	}
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	userID, err := x.authService.ValidateAccessToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
		return
	}

	utils.SetGinContextUserID(c, userID)
	utils.SetGinContextAccessToken(c, accessToken)
	c.Next()
}
