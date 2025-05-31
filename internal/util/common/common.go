package common

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"learn-fiber/internal/constant"
	"learn-fiber/internal/ierror"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAccessToken(c *fiber.Ctx) (token string, err error) {
	authorization := c.Get(fiber.HeaderAuthorization)
	if authorization == "" {
		err = errors.New("no authorization header")
		return
	}
	values := strings.Split(authorization, " ")
	if len(values) != 2 {
		err = errors.New("authorization header malformed")
		return
	}
	token = values[1]
	return
}

func GetRequestBody[T any](c *fiber.Ctx) (T, error) {
	var body T

	value := c.Locals(constant.ContextValidatedBody)
	if value == nil {
		err := c.BodyParser(&body)
		if err != nil {
			return body, ierror.ErrInvalidRequestBody
		}
		return body, nil
	}

	v, ok := value.(T)
	if !ok {
		return body, errors.New("request body parsing invalid type assertion")
	}

	return v, nil
}

func GetQueryParam[T any](c *fiber.Ctx) (T, error) {
	var query T

	value := c.Locals(constant.ContextValidatedBody)
	if value == nil {
		err := c.QueryParser(&query)
		if err != nil {
			return query, ierror.ErrInvalidRequestBody
		}
		return query, nil
	}

	v, ok := value.(T)
	if !ok {
		return query, errors.New("request query parsing invalid type assertion")
	}

	return v, nil
}

func GetUser(c *fiber.Ctx) string {
	claims := c.Locals(constant.ContextUser).(jwt.MapClaims)
	return claims["username"].(string)
}

func SplitVersion(version string) (int, int, error) {
	parts := strings.SplitN(version, ".", 2)
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	if len(parts) == 1 {
		afterSplit := fmt.Sprintf("%d.%d", major, 0)
		if afterSplit != version {
			return 0, 0, errors.New("invalid version")
		}
		return major, 0, nil // Only major version present
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}
	afterSplit := fmt.Sprintf("%d.%d", major, minor)
	if afterSplit != version {
		return 0, 0, errors.New("invalid version")
	}

	return major, minor, nil // Major and minor versions
}

func CheckValidVersion(major, minor, maxMajor, maxMinor int) error {
	if major <= maxMajor {
		if major == maxMajor {
			if minor <= maxMinor {
				return ierror.NewClientError(http.StatusBadRequest, ierror.ErrCodeInvalidVersion, "invalid minor version")
			}
		} else {
			return ierror.NewClientError(http.StatusBadRequest, ierror.ErrCodeInvalidVersion, "invalid major version")
		}
	}
	return nil
}

func UniqueValues(list []string) []string {
	keys := make(map[string]bool)
	result := []string{}
	for _, entry := range list {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}

func PointerInt(v int) *int {
	return &v
}

func TimePtr(t time.Time) *time.Time {
	return &t
}

func Float64Ptr(v float64) *float64 {
	return &v
}

func BoolPtr(v bool) *bool {
	return &v
}
