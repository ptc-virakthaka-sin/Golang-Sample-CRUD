package middleware

import (
	"learn-fiber/config"
	"learn-fiber/pkg/redis"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Limiter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := "limiter:" + c.IP()
		if count, _ := redis.Incr(key); int(count) > config.Cfg.RateLimit {
			return c.SendStatus(http.StatusTooManyRequests)
		} else if count == 1 {
			redis.Exp(key, time.Minute)
		}
		return c.Next()
	}
}
