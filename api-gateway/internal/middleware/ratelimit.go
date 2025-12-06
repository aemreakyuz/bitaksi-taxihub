package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	requests map[string][]time.Time
	mu       sync.Mutex
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		rl.mu.Lock()
		defer rl.mu.Unlock()

		now := time.Now()

		if timestamps, exists := rl.requests[ip]; exists {
			var validRequests []time.Time
			for _, t := range timestamps {
				if now.Sub(t) < rl.window {
					validRequests = append(validRequests, t)
				}
			}
			rl.requests[ip] = validRequests
		}

		if len(rl.requests[ip]) >= rl.limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		rl.requests[ip] = append(rl.requests[ip], now)

		c.Next()
	}
}
