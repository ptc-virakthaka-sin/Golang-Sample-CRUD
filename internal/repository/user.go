package repository

import (
	"learn-fiber/internal/model"

	"gorm.io/gorm"
)

type User interface {
	GetUser(username string) (model.User, error)
	UpdatePass(username, password string) error
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return &user{db: db}
}

func (h *user) GetUser(username string) (entity model.User, err error) {
	return entity, h.db.First(&entity, "username = ?", username).Error
}

func (h *user) UpdatePass(username, password string) error {
	return h.db.Model(&model.User{}).Where("username = ?", username).
		Update("password", password).Error
}
