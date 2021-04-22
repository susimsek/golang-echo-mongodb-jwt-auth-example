package exception

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ResourceNotFoundException(resourceName string, fieldName string, fieldValue string) error {
	msg := fmt.Sprintf("%s not found with %s : %s", resourceName, fieldName, fieldValue)
	return echo.NewHTTPError(http.StatusNotFound, msg)
}

func BadRequestException(msg string) error {
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func ConflictException(resourceName string, fieldName string, fieldValue string) error {
	msg := fmt.Sprintf("%s with %s : %s already exists", resourceName, fieldName, fieldValue)
	return echo.NewHTTPError(http.StatusConflict, msg)
}

func UnauthorizedException() error {
	return echo.ErrUnauthorized
}
