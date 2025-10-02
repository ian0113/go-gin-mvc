package middlewares

import (
	"time"

	"github.com/ian0113/go-gin-mvc/infra"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiMiddleware struct {
	logger *zap.Logger
}

func NewApiMiddleware() *ApiMiddleware {
	return &ApiMiddleware{
		logger: infra.GetLogger().Named("api.middleware"),
	}
}

func (x *ApiMiddleware) Logger(c *gin.Context) {
	start := time.Now()
	c.Next()
	x.logger.Sugar().Infof("[%s] %s %d %s",
		c.Request.Method,
		c.Request.URL.Path,
		c.Writer.Status(),
		time.Since(start),
	)
}
