package ierror

import (
	"errors"
	"learn-fiber/api/response"
	"learn-fiber/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HandleErrorResponse() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		message := err.Error()
		logger.L.Error(message)

		var status int
		var code string
		var traceID string
		var validationErrors []response.ValidationError

		reqId := c.Locals(requestid.ConfigDefault.ContextKey)
		if reqId != nil {
			traceID = reqId.(string)
		}

		var fiberErr *fiber.Error
		var authenticateErr *AuthenticationError
		var authorizeErr *AuthorizationError
		var validateErr *ValidationError
		var clientErr *ClientError
		var severErr *ServerError

		switch {
		case errors.As(err, &fiberErr):
			status = fiberErr.Code
			code = strconv.Itoa(fiberErr.Code)
		case errors.As(err, &authenticateErr):
			status = authenticateErr.Status
			code = authenticateErr.Code
		case errors.As(err, &authorizeErr):
			status = authorizeErr.Status
			code = authorizeErr.Code
		case errors.As(err, &clientErr):
			status = clientErr.Status
			code = clientErr.Code
		case errors.As(err, &severErr):
			status = severErr.Status
			code = severErr.Code
		case errors.As(err, &validateErr):
			status = validateErr.Status
			code = validateErr.Code
			validationErrors = validateErr.Errors
		case errors.Is(err, gorm.ErrRecordNotFound),
			errors.Is(err, gorm.ErrInvalidTransaction),
			errors.Is(err, gorm.ErrNotImplemented),
			errors.Is(err, gorm.ErrMissingWhereClause),
			errors.Is(err, gorm.ErrUnsupportedRelation),
			errors.Is(err, gorm.ErrPrimaryKeyRequired),
			errors.Is(err, gorm.ErrModelValueRequired),
			errors.Is(err, gorm.ErrModelAccessibleFieldsRequired),
			errors.Is(err, gorm.ErrSubQueryRequired),
			errors.Is(err, gorm.ErrInvalidData),
			errors.Is(err, gorm.ErrUnsupportedDriver),
			errors.Is(err, gorm.ErrRegistered),
			errors.Is(err, gorm.ErrInvalidField),
			errors.Is(err, gorm.ErrEmptySlice),
			errors.Is(err, gorm.ErrDryRunModeUnsupported),
			errors.Is(err, gorm.ErrInvalidDB),
			errors.Is(err, gorm.ErrInvalidValue),
			errors.Is(err, gorm.ErrInvalidValueOfLength),
			errors.Is(err, gorm.ErrPreloadNotAllowed),
			errors.Is(err, gorm.ErrDuplicatedKey),
			errors.Is(err, gorm.ErrForeignKeyViolated),
			errors.Is(err, gorm.ErrCheckConstraintViolated):

			status = http.StatusInternalServerError
			code = ErrCodeDatabaseError

			if errors.Is(err, gorm.ErrRecordNotFound) {
				message = "record not found"
			} else {
				message = "database error cannot process this request at the moment. Please try again later."
			}
		default:
			status = http.StatusInternalServerError
			code = ErrCodeGeneralError
			message = "something went wrong"
		}

		return c.Status(status).
			JSON(response.ErrorResponse{
				Error: response.ErrorObject{
					Code:             code,
					Message:          message,
					ValidationErrors: validationErrors,
					TraceID:          traceID,
				},
			})
	}
}
