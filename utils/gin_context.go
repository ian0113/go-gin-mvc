package utils

import "github.com/gin-gonic/gin"

func SetGinContextUserID(c *gin.Context, userID uint) {
	c.Set("user_id", userID)
}

func GetGinContextUserID(c *gin.Context) (uint, bool) {
	userIDInf, ok1 := c.Get("user_id")
	userID, ok2 := userIDInf.(uint)
	if !ok1 || !ok2 {
		return 0, false
	}
	return userID, true
}

func SetGinContextAccessToken(c *gin.Context, token string) {
	c.Set("access_token", token)
}

func GetGinContextAccessToken(c *gin.Context) (string, bool) {
	accessTokenInf, ok1 := c.Get("access_token")
	accessToken, ok2 := accessTokenInf.(string)
	if !ok1 || !ok2 {
		return "", false
	}
	return accessToken, true
}
