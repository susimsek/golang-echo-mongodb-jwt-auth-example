package security

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config2 "golang-echo-mongodb-jwt-auth-example/config"
	"golang-echo-mongodb-jwt-auth-example/model"
)

var whiteListPaths = []string{
	"/favicon.ico",
	"/api",
	"/api/*",
	"/api/v1/login",
	"/api/v1/signup",
}

// change default error message
func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func WebSecurityConfig(e *echo.Echo) {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(config2.JWTSecret),
		Skipper:    skipAuth,
	}
	e.Use(middleware.JWTWithConfig(config))
}

func skipAuth(e echo.Context) bool {
	// Skip authentication for and signup login requests
	for _, path := range whiteListPaths {
		if path == e.Path() {
			return true
		}
	}
	return false
}
