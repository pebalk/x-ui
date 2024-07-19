package middleware

import (
	"net"
	"net/http"
	"strings"

	"x-ui/logger"

	"github.com/gin-gonic/gin"
)

func DomainValidatorMiddleware(domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		if colonIndex := strings.LastIndex(host, ":"); colonIndex != -1 {
			host, _, _ = net.SplitHostPort(c.Request.Host)
		}
		logger.Info("!!!!----host is ", host)
		logger.Info("!!!!----domain is ", domain)

		if host != domain {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
