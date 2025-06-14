package dto

import (
	"learn-fiber/pkg/database/pagination"
)

type QueryParams struct {
	Query string `json:"query"`
	pagination.PageRequest
}
