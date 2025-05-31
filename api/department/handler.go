package department

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
	svc service.Department
}

func NewHandler(db *gorm.DB) *Handler {
	svc := service.NewDepartment(db)
	return &Handler{svc: svc}
}

func (h *Handler) GetList(c *fiber.Ctx) error {
	query, _ := common.GetQueryParam[dto.QueryParams](c)
	if hasError, err := validator.V.Valid(query); hasError {
		return ierror.NewValidationError(err)
	}
	result, page, err := h.svc.GetAllDepartment(query)
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
	result, err := h.svc.GetDepartment(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.SendSuccessResponse(c, err.Error())
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
	req, _ := common.GetRequestBody[dto.DepartmentCreateRequest](c)
	if hasError, err := validator.V.Valid(req); hasError {
		return ierror.NewValidationError(err)
	}
	result, err := h.svc.CreateDepartment(req)
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
	req, _ := common.GetRequestBody[dto.DepartmentUpdateRequest](c)
	if hasError, err := validator.V.Valid(req); hasError {
		return ierror.NewValidationError(err)
	}
	result, err := h.svc.UpdateDepartment(req)
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
	if err := h.svc.DeleteDepartment(id); err != nil {
		return ierror.NewServerError(ierror.ErrCodeDatabaseError, err.Error())
	}
	return response.SendSuccessResponse(c, constant.Success)
}
