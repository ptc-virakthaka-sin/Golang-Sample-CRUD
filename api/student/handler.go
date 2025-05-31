package student

import (
	"errors"
	"learn-fiber/api/response"
	"learn-fiber/internal/constant"
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
	svc service.Student
}

func NewHandler(db *gorm.DB) *Handler {
	svc := service.NewStudent(db)
	return &Handler{svc: svc}
}

func (h *Handler) GetList(c *fiber.Ctx) error {
	query, _ := common.GetQueryParam[dto.QueryParams](c)
	if hasError, err := validator.V.Valid(query); hasError {
		return ierror.NewValidationError(err)
	}
	result, page, err := h.svc.GetAllStudent(query)
	if err != nil {
		return ierror.NewServerError(ierror.ErrCodeDatabaseError, err.Error())
	}
	var res []Response
	if err = copier.Copy(&res, &result); err != nil {
		return ierror.NewServerError(ierror.ErrCodeDtoError, err.Error())
	}
	return response.SendSuccessResponsePaging(c, res, page)
}

func (h *Handler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.svc.GetStudent(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ierror.NewClientError(200, ierror.ErrCodeDataNotFound, err.Error())
		}
		return ierror.NewServerError(ierror.ErrCodeDatabaseError, err.Error())
	}
	var res Response
	if err = copier.Copy(&res, &result); err != nil {
		return ierror.NewServerError(ierror.ErrCodeDtoError, err.Error())
	}
	return response.SendSuccessResponse(c, res)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	req, err := common.GetRequestBody[dto.StudentCreateRequest](c)
	if err != nil {
		return err
	}
	result, err := h.svc.CreateStudent(req)
	if err != nil {
		return ierror.NewServerError(ierror.ErrCodeDatabaseError, err.Error())
	}
	var res Response
	if err = copier.Copy(&res, &result); err != nil {
		return ierror.NewServerError(ierror.ErrCodeDtoError, err.Error())
	}
	return response.SendSuccessResponse(c, res)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	req, err := common.GetRequestBody[dto.StudentUpdateRequest](c)
	if err != nil {
		return err
	}
	result, err := h.svc.UpdateStudent(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ierror.NewClientError(200, ierror.ErrCodeDataNotFound, err.Error())
		}
		return ierror.NewServerError(ierror.ErrCodeDatabaseError, err.Error())
	}
	var res Response
	if err = copier.Copy(&res, &result); err != nil {
		return ierror.NewServerError(ierror.ErrCodeDtoError, err.Error())
	}
	return response.SendSuccessResponse(c, res)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.DeleteStudent(id); err != nil {
		return ierror.NewServerError(ierror.ErrCodeDatabaseError, err.Error())
	}
	return response.SendSuccessResponse(c, constant.Success)
}
