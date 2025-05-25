package response

import (
	"learn-fiber/internal/constant"
	"learn-fiber/pkg/database/pagination"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Message string                   `json:"message,omitempty"`
	Paging  *pagination.PageResponse `json:"pagination,omitempty"`
	Data    interface{}              `json:"data,omitempty"`
}

func SendSuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.
		Status(http.StatusOK).
		JSON(SuccessResponse{
			Message: constant.Success,
			Data:    data,
		})
}

func SendFailResponse(c *fiber.Ctx, data interface{}) error {
	return c.
		Status(http.StatusOK).
		JSON(SuccessResponse{
			Message: constant.Error,
			Data:    data,
		})
}

func SendSuccessResponsePaging(c *fiber.Ctx, data interface{}, page pagination.PageResponse) error {
	return c.
		Status(http.StatusOK).
		JSON(SuccessResponse{
			Message: constant.Success,
			Paging:  &page,
			Data:    data,
		})
}
