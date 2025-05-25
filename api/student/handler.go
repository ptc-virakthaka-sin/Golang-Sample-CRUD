package student

import (
	"errors"
	"learn-fiber/api/response"
	"learn-fiber/internal/dto"
	"learn-fiber/internal/ierror"
	"learn-fiber/internal/service"
	"learn-fiber/internal/util/common"
	"learn-fiber/internal/util/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Handler struct {
	service service.Student
}

func NewHandler(db *gorm.DB) *Handler {
	svc := service.NewStudent(db)
	return &Handler{service: svc}
}

func (h *Handler) GetList(c *fiber.Ctx) error {
	query, _ := common.GetQueryParam[dto.QueryParams](c)
	if hasError, err := validator.V.Valid(query); hasError {
		return ierror.NewValidationError(err)
	}
	result, page, err := h.service.GetAllStudent(query)
	if err != nil {
		return response.SendFailResponse(c, err.Error())
	}
	var res []Response
	_ = copier.Copy(&res, &result)
	return response.SendSuccessResponsePaging(c, res, page)
}

func (h *Handler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.GetStudent(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.SendSuccessResponse(c, nil)
		}
		return err
	}
	var res Response
	_ = copier.Copy(&res, &result)
	return response.SendSuccessResponse(c, res)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	req, _ := common.GetRequestBody[dto.StudentCreateRequest](c)
	if hasError, err := validator.V.Valid(req); hasError {
		return ierror.NewValidationError(err)
	}
	result, err := h.service.CreateStudent(req)
	if err != nil {
		return response.SendFailResponse(c, err.Error())
	}
	var res Response
	_ = copier.Copy(&res, &result)
	return response.SendSuccessResponse(c, res)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	req, _ := common.GetRequestBody[dto.StudentUpdateRequest](c)
	if hasError, err := validator.V.Valid(req); hasError {
		return ierror.NewValidationError(err)
	}
	result, err := h.service.UpdateStudent(req)
	if err != nil {
		return response.SendFailResponse(c, err.Error())
	}
	var res Response
	_ = copier.Copy(&res, &result)
	return response.SendSuccessResponse(c, res)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	_ = h.service.DeleteStudent(id)
	return response.SendSuccessResponse(c, nil)
}
