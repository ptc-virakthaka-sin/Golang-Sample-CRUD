package student

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddRoutes(r fiber.Router, db *gorm.DB) {
	handler := NewHandler(db)
	router := r.Group("/student")
	router.Put("", handler.Update)
	router.Get("", handler.GetList)
	router.Post("", handler.Create)
	router.Get(":id", handler.GetById)
	router.Delete(":id", handler.Delete)
}
