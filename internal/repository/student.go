package repository

import (
	"learn-fiber/internal/dto"
	"learn-fiber/internal/model"
	"learn-fiber/pkg/database/pagination"
	"learn-fiber/pkg/database/util"

	"gorm.io/gorm"
)

type Student interface {
	GetAllStudent(query dto.QueryParams) ([]model.Student, pagination.PageResponse, error)
	CreateStudent(model *model.Student) (*model.Student, error)
	UpdateStudent(model *model.Student) (*model.Student, error)
	GetStudent(Id string) (*model.Student, error)
	DeleteStudent(Id string) error
}

func NewStudent(db *gorm.DB) Student {
	return &Handler{db: db}
}

func (h *Handler) GetAllStudent(query dto.QueryParams) (result []model.Student, page pagination.PageResponse, err error) {
	db := h.db.Model(&result)
	if query.Query != "" {
		query := util.EscapeLike("%", "%", query.Query)
		db = db.Where("code LIKE ? OR name_en LIKE ? OR name_km LIKE ?", query, query, query)
	}
	db.Count(&page.Total)
	page = pagination.GetPageResponse(query.Page, query.Limit, page.Total)
	paginate := pagination.ToPaginate(query.PageRequest)
	db = util.WherePaginateAndOrderBy(db, paginate)
	db.Find(&result)
	return
}

func (h *Handler) CreateStudent(model *model.Student) (*model.Student, error) {
	return model, h.db.Create(model).Error
}

func (h *Handler) UpdateStudent(model *model.Student) (*model.Student, error) {
	return model, h.db.Save(model).Error
}

func (h *Handler) GetStudent(Id string) (model *model.Student, err error) {
	return model, h.db.Find(&model, "id = ?", Id).Error
}

func (h *Handler) DeleteStudent(Id string) error {
	return h.db.Delete(&model.Student{Id: Id}).Error
}
