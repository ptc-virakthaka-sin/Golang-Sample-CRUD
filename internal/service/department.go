package service

import (
	"learn-fiber/internal/dto"
	"learn-fiber/internal/model"
	"learn-fiber/internal/repository"
	"learn-fiber/pkg/database/pagination"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Department interface {
	GetAllDepartment(query dto.QueryParams) ([]model.Department, pagination.PageResponse, error)
	CreateDepartment(req dto.DepartmentCreateRequest) (*model.Department, error)
	UpdateDepartment(req dto.DepartmentUpdateRequest) (*model.Department, error)
	GetDepartment(Id string) (*model.Department, error)
	DeleteDepartment(Id string) error
}

type department struct {
	repo repository.Department
}

func NewDepartment(db *gorm.DB) Department {
	repo := repository.NewDepartment(db)
	return &department{repo: repo}
}

func (h *department) GetAllDepartment(query dto.QueryParams) ([]model.Department, pagination.PageResponse, error) {
	return h.repo.GetAllDepartment(query)
}

func (h *department) CreateDepartment(req dto.DepartmentCreateRequest) (entity *model.Department, err error) {
	_ = copier.Copy(entity, &req)
	return h.repo.CreateDepartment(entity)
}

func (h *department) UpdateDepartment(req dto.DepartmentUpdateRequest) (entity *model.Department, err error) {
	_ = copier.Copy(entity, &req)
	return h.repo.UpdateDepartment(entity)
}

func (h *department) GetDepartment(Id string) (*model.Department, error) {
	return h.repo.GetDepartment(Id)
}

func (h *department) DeleteDepartment(Id string) error {
	return h.repo.DeleteDepartment(Id)
}
