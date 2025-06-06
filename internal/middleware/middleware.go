package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"learn-fiber/config"
	"learn-fiber/internal/constant"
	"learn-fiber/internal/ierror"
	"slices"
	"strings"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get(fiber.HeaderAuthorization)
		if !strings.HasPrefix(auth, constant.Bearer) {
			return ierror.NewAuthenticationError(ierror.ErrCodeAuthenticationError, "Invalid token")
		}
		tokenStr := strings.TrimPrefix(auth, constant.Bearer)
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Cfg.JWT.Secret), nil
		})
		if err != nil || !token.Valid {
			return ierror.NewAuthenticationError(ierror.ErrCodeAuthenticationError, "Invalid token")
		}
		c.Locals(constant.ContextUser, token.Claims.(jwt.MapClaims))
		return c.Next()
	}
}

func RequireRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals(constant.ContextUser).(jwt.MapClaims)
		if r, ok := claims["role"].(string); ok && slices.Contains(roles, r) {
			return c.Next()
		}
		return ierror.NewAuthorizationError(ierror.ErrCodeAuthorizationError, "Insufficient permissions")
	}
}
