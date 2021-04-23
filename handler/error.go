package handler

import (
	"github.com/labstack/echo/v4"
	"golang-echo-mongodb-jwt-auth-example/util"
	"net/http"
	"time"
)

// APIError example
type APIError struct {
	Status    int    `json:"status" xml:"status"`
	Message   string `json:"message" xml:"message"`
	Path      string `json:"path" xml:"path"`
	Timestamp int64  `json:"timestamp" xml:"timestamp"`
}

func ErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	code := he.Code
	message := he.Message
	if m, ok := he.Message.(string); ok {
		message = &APIError{
			Status:    code,
			Message:   m,
			Path:      c.Request().RequestURI,
			Timestamp: time.Now().Unix(),
		}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			util.Negotiate(c, code, message)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
