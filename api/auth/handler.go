package auth

import (
	"learn-fiber/api/response"
	"learn-fiber/config"
	"learn-fiber/internal/util/common"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	req, err := common.GetRequestBody[Request](c)
	if err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	secret := []byte(config.Cfg.JWT.Secret)
	res, err := token.SignedString(secret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
	}
	return response.SendSuccessResponse(c, res)
}
