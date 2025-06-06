package repository

import (
	"learn-fiber/internal/model"

	"gorm.io/gorm"
)

type Token interface {
	Remove(token string) error
	Add(model.Token) (model.Token, error)
	Get(refresh string) (model.Token, error)
	Find(ip, username string) (model.Token, error)
	Renew(entity model.Token) (model.Token, error)
}

type token struct {
	db *gorm.DB
}

func NewToken(db *gorm.DB) Token {
	return &token{db: db}
}

func (h *token) Remove(token string) error {
	return h.db.Delete(&model.Token{AccessToken: token}).Error
}

func (h *token) Add(entity model.Token) (model.Token, error) {
	return entity, h.db.Create(&entity).Error
}

func (h *token) Get(refresh string) (entity model.Token, err error) {
	return entity, h.db.First(&entity, "refresh_token = ?", refresh).Error
}

func (h *token) Find(ip, username string) (entity model.Token, err error) {
	return entity, h.db.First(&entity, "ip = ? and username = ?", ip, username).Error
}

func (h *token) Renew(entity model.Token) (model.Token, error) {
	return entity, h.db.Updates(entity).Error
}
