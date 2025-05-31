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
	CreateStudent(entity model.Student) (model.Student, error)
	UpdateStudent(entity model.Student) (model.Student, error)
	GetStudent(Id string) (model.Student, error)
	DeleteStudent(Id string) error
}

type student struct {
	db *gorm.DB
}

func NewStudent(db *gorm.DB) Student {
	return &student{db: db}
}

func (h *student) GetAllStudent(query dto.QueryParams) (result []model.Student, page pagination.PageResponse, err error) {
	db := h.db.Model(&result)
	if query.Query != "" {
		search := util.EscapeLike("%", "%", query.Query)
		db = db.Where("code LIKE ? OR name_en LIKE ? OR name_km LIKE ?", search, search, search)
	}
	db.Count(&page.Total)
	page = pagination.GetPageResponse(query.Page, query.Limit, page.Total)
	paginate := pagination.ToPaginate(query.PageRequest)
	db = util.WherePaginateAndOrderBy(db, paginate)
	err = db.Find(&result).Error
	return
}

func (h *student) CreateStudent(entity model.Student) (model.Student, error) {
	return entity, h.db.Create(&entity).Error
}

func (h *student) UpdateStudent(entity model.Student) (model.Student, error) {
	existing, err := h.GetStudent(entity.Id)
	if err != nil {
		return existing, err
	}
	return existing, h.db.Model(&existing).Updates(entity).Error
}

func (h *student) GetStudent(Id string) (entity model.Student, err error) {
	return entity, h.db.First(&entity, "id = ?", Id).Error
}

func (h *student) DeleteStudent(Id string) error {
	return h.db.Delete(&model.Student{Id: Id}).Error
}
