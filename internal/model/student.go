package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	Id     string `gorm:"column:id; type:varchar(36); not null; primaryKey"`
	Code   string `gorm:"column:code; type:varchar(255)"`
	NameEN string `gorm:"column:name_en; type:varchar(255)"`
	NameKM string `gorm:"column:name_km; type:varchar(255)"`
}

func (Student) TableName() string {
	return "students"
}

func (m *Student) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.NewString()
	return
}
