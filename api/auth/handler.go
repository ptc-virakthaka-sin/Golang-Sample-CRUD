package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"learn-fiber/api/response"
	"learn-fiber/internal/constant"
	"learn-fiber/internal/dto"
	"learn-fiber/internal/ierror"
	"learn-fiber/internal/service"
	"learn-fiber/internal/util/common"
	"learn-fiber/internal/util/validator"
)

type Handler struct {
	svc service.Auth
}

func NewHandler(db *gorm.DB) *Handler {
	svc := service.NewAuth(db)
	return &Handler{svc: svc}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	req, _ := common.GetRequestBody[dto.LoginRequest](c)
	if hasError, err := validator.V.Valid(req); hasError {
		return ierror.NewValidationError(err)
	}
	result, err := h.svc.Login(req)
	if err != nil {
		return ierror.NewClientError(200, ierror.ErrCodeAuthenticationError, err.Error())
	}
	var res TokenResponse
	if err = copier.Copy(&res, &result); err != nil {
		return ierror.NewServerError(ierror.ErrCodeDtoError, err.Error())
	}
	return response.SendSuccessResponse(c, res)
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	token, err := common.GetAccessToken(c)
	if err != nil {
		return ierror.NewAuthenticationError(ierror.ErrCodeAuthenticationError, err.Error())
	}
	if err := h.svc.Logout(token); err != nil {
		return ierror.NewServerError(ierror.ErrCodeTokenError, err.Error())
	}
	return response.SendSuccessResponse(c, constant.Success)
}

func (h *Handler) RenewToken(c *fiber.Ctx) error {
	req, _ := common.GetRequestBody[dto.RenewTokenRequest](c)
	if hasError, err := validator.V.Valid(req); hasError {
		return ierror.NewValidationError(err)
	}
	result, err := h.svc.RenewToken(req)
	if err != nil {
		return ierror.NewClientError(200, ierror.ErrCodeAuthenticationError, err.Error())
	}
	var res TokenResponse
	if err = copier.Copy(&res, &result); err != nil {
		return ierror.NewServerError(ierror.ErrCodeDtoError, err.Error())
	}
	return response.SendSuccessResponse(c, res)

}

func (h *Handler) Info(c *fiber.Ctx) error {
	result, err := h.svc.Info(common.GetUser(c))
	if err != nil {
		return ierror.NewClientError(200, ierror.ErrCodeAuthenticationError, err.Error())
	}
	var res UserResponse
	if err = copier.Copy(&res, &result); err != nil {
		return ierror.NewServerError(ierror.ErrCodeDtoError, err.Error())
	}
	return response.SendSuccessResponse(c, res)
}

func (h *Handler) ChangePassword(c *fiber.Ctx) error {
	req, _ := common.GetRequestBody[dto.ChangePassRequest](c)
	if hasError, err := validator.V.Valid(req); hasError {
		return ierror.NewValidationError(err)
	}
	if err := h.svc.ChangePassword(common.GetUser(c), req); err != nil {
		return ierror.NewClientError(200, ierror.ErrCodeValidationError, err.Error())
	}
	return response.SendSuccessResponse(c, constant.Success)
}
