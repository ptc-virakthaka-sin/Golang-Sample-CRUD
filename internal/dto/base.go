package dto

import (
	"learn-fiber/pkg/database/pagination"
)

type QueryParams struct {
	Query string `json:"query" validate:"required"`
	pagination.PageRequest
}
