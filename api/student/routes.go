package student

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"learn-fiber/internal/middleware"
)

func AddRoutes(r fiber.Router, db *gorm.DB) {
	handler := NewHandler(db)
	router := r.Group("/student",
		middleware.JWTProtected(),
		middleware.RequireRole("admin"),
	)
	router.Put("", handler.Update)
	router.Get("", handler.GetList)
	router.Post("", handler.Create)
	router.Get(":id", handler.GetById)
	router.Delete(":id", handler.Delete)
}
