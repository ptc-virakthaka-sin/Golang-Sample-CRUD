package dto

type DepartmentCreateRequest struct {
	Code         string `json:"code" validate:"required,min=3,max=100"`
	DepartmentEN string `json:"departmentEN" validate:"required"`
	DepartmentKM string `json:"departmentKM" validate:"required"`
}

type DepartmentUpdateRequest struct {
	Id           string `json:"id" validate:"required"`
	Code         string `json:"code" validate:"required,min=3,max=100"`
	DepartmentEN string `json:"departmentEN" validate:"required"`
	DepartmentKM string `json:"departmentKM" validate:"required"`
}
