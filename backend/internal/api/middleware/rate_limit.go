package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/neurowatt/aiwatt-backend/pkg/response"
)

// RateLimit returns a middleware that enforces per-IP request rate limits using Redis.
// limit is the maximum number of requests allowed per window.
// window is the duration of the sliding window.
func RateLimit(rdb *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("rl:ip:%s", ip)

		ctx := context.Background()
		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			// Redis failure: fail open (don't block requests)
			c.Next()
			return
		}
		if count == 1 {
			rdb.Expire(ctx, key, window)
		}
		if count > int64(limit) {
			msg := fmt.Sprintf("rate limit exceeded: max %d requests per %s", limit, window)
			c.Header("Retry-After", fmt.Sprintf("%.0f", window.Seconds()))
			c.JSON(http.StatusTooManyRequests, response.Envelope{
				Success: false,
				Data:    nil,
				Error:   &msg,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
