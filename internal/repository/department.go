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
	CreateDepartment(model *model.Department) (*model.Department, error)
	UpdateDepartment(model *model.Department) (*model.Department, error)
	GetDepartment(Id string) (*model.Department, error)
	DeleteDepartment(Id string) error
}

func NewDepartment(db *gorm.DB) Department {
	return &Handler{db: db}
}

func (h *Handler) GetAllDepartment(query dto.QueryParams) (result []model.Department, page pagination.PageResponse, err error) {
	db := h.db.Model(&result)
	if query.Query != "" {
		query := util.EscapeLike("%", "%", query.Query)
		db = db.Where("code LIKE ? OR department_en LIKE ? OR department_km LIKE ?", query, query, query)
	}
	db.Count(&page.Total)
	page = pagination.GetPageResponse(query.Page, query.Limit, page.Total)
	paginate := pagination.ToPaginate(query.PageRequest)
	db = util.WherePaginateAndOrderBy(db, paginate)
	db.Find(&result)
	return
}

func (h *Handler) CreateDepartment(model *model.Department) (*model.Department, error) {
	return model, h.db.Create(model).Error
}

func (h *Handler) UpdateDepartment(model *model.Department) (*model.Department, error) {
	return model, h.db.Save(model).Error
}

func (h *Handler) GetDepartment(Id string) (model *model.Department, err error) {
	return model, h.db.Find(&model, "id = ?", Id).Error
}

func (h *Handler) DeleteDepartment(Id string) error {
	return h.db.Delete(&model.Department{Id: Id}).Error
}
