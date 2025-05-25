package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddRoutes(r fiber.Router, db *gorm.DB) {
	handler := NewHandler(db)
	router := r.Group("/auth")
	router.Post("login", handler.Login)
}
