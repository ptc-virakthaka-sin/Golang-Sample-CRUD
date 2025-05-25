package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	Id           string `gorm:"column:id; type:varchar(36); not null; primaryKey"`
	Code         string `gorm:"column:code; type:varchar(255)"`
	DepartmentEN string `gorm:"column:department_en; type:varchar(255)"`
	DepartmentKM string `gorm:"column:department_km; type:varchar(255)"`
}

func (Department) TableName() string {
	return "departments"
}

func (m *Department) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.NewString()
	return
}
