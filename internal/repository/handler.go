package repository

import (
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}
