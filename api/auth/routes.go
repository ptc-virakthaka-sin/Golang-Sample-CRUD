package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"learn-fiber/internal/middleware"
)

func AddRoutes(r fiber.Router, db *gorm.DB) {
	handler := NewHandler(db)
	router := r.Group("/auth")
	router.Post("login", handler.Login)
	router.Post("renew-token", handler.RenewToken)
	router.Get("info", middleware.JWTProtected(), handler.Info)
	router.Post("logout", middleware.JWTProtected(), handler.Logout)
	router.Post("change-password", middleware.JWTProtected(), handler.ChangePassword)
}
