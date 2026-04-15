package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORS returns a middleware that sets permissive CORS headers for the given allowed origins.
// Pass allowedOrigins as a comma-separated list or "*" for dev.
func CORS(allowedOrigins string) gin.HandlerFunc {
	origins := strings.Split(allowedOrigins, ",")
	originSet := make(map[string]bool, len(origins))
	for _, o := range origins {
		originSet[strings.TrimSpace(o)] = true
	}
	allowAll := allowedOrigins == "*"

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		if allowAll || originSet[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
		}
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type,X-HMAC-Signature")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
