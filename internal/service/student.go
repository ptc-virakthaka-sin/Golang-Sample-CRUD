package service

import (
	"learn-fiber/internal/dto"
	"learn-fiber/internal/model"
	"learn-fiber/internal/repository"
	"learn-fiber/pkg/database/pagination"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Student interface {
	GetAllStudent(query dto.QueryParams) ([]model.Student, pagination.PageResponse, error)
	CreateStudent(req dto.StudentCreateRequest) (*model.Student, error)
	UpdateStudent(req dto.StudentUpdateRequest) (*model.Student, error)
	GetStudent(Id string) (*model.Student, error)
	DeleteStudent(Id string) error
}

type student struct {
	repo repository.Student
}

func NewStudent(db *gorm.DB) Student {
	repo := repository.NewStudent(db)
	return &student{repo: repo}
}

func (h *student) GetAllStudent(query dto.QueryParams) ([]model.Student, pagination.PageResponse, error) {
	return h.repo.GetAllStudent(query)
}

func (h *student) CreateStudent(req dto.StudentCreateRequest) (entity *model.Student, err error) {
	_ = copier.Copy(&entity, &req)
	return h.repo.CreateStudent(entity)
}

func (h *student) UpdateStudent(req dto.StudentUpdateRequest) (entity *model.Student, err error) {
	_ = copier.Copy(&entity, &req)
	return h.repo.UpdateStudent(entity)
}

func (h *student) GetStudent(Id string) (*model.Student, error) {
	return h.repo.GetStudent(Id)
}

func (h *student) DeleteStudent(Id string) error {
	return h.repo.DeleteStudent(Id)
}
