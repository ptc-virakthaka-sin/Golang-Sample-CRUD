package repository

import (
	"learn-fiber/internal/dto"
	"learn-fiber/internal/model"
	"learn-fiber/pkg/database/pagination"
	"learn-fiber/pkg/database/util"

	"gorm.io/gorm"
)

type Department interface {
	GetAllDepartment(query dto.QueryParams) ([]model.Department, pagination.PageResponse, error)
	CreateDepartment(entity model.Department) (model.Department, error)
	UpdateDepartment(entity model.Department) (model.Department, error)
	GetDepartment(Id string) (model.Department, error)
	DeleteDepartment(Id string) error
}

type department struct {
	db *gorm.DB
}

func NewDepartment(db *gorm.DB) Department {
	return &department{db: db}
}

func (h *department) GetAllDepartment(query dto.QueryParams) (result []model.Department, page pagination.PageResponse, err error) {
	db := h.db.Model(&result)
	if query.Query != "" {
		search := util.EscapeLike("%", "%", query.Query)
		db = db.Where("code LIKE ? OR department_en LIKE ? OR department_km LIKE ?", search, search, search)
	}
	db.Count(&page.Total)
	page = pagination.GetPageResponse(query.Page, query.Limit, page.Total)
	paginate := pagination.ToPaginate(query.PageRequest)
	db = util.WherePaginateAndOrderBy(db, paginate)
	err = db.Find(&result).Error
	return
}

func (h *department) CreateDepartment(entity model.Department) (model.Department, error) {
	return entity, h.db.Create(&entity).Error
}

func (h *department) UpdateDepartment(entity model.Department) (model.Department, error) {
	existing, err := h.GetDepartment(entity.Id)
	if err != nil {
		return existing, err
	}
	return existing, h.db.Model(&existing).Updates(entity).Error
}

func (h *department) GetDepartment(Id string) (entity model.Department, err error) {
	return entity, h.db.First(&entity, "id = ?", Id).Error
}

func (h *department) DeleteDepartment(Id string) error {
	return h.db.Delete(&model.Department{Id: Id}).Error
}
