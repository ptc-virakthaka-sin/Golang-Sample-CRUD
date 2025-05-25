package dto

type StudentCreateRequest struct {
	Code   string `json:"code" validate:"required"`
	NameEN string `json:"nameEN" validate:"required"`
	NameKM string `json:"nameKM" validate:"required"`
}

type StudentUpdateRequest struct {
	Id     string `json:"id" validate:"required"`
	Code   string `json:"code" validate:"required"`
	NameEN string `json:"nameEN" validate:"required"`
	NameKM string `json:"nameKM" validate:"required"`
}
