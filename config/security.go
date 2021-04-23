package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang-echo-mongodb-jwt-auth-example/model"
)

// change default error message
func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func WebSecurityConfig(e *echo.Echo) {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(JWTSecret),
		Skipper:    skipAuth,
	}
	e.Use(middleware.JWTWithConfig(config))
}

func skipAuth(e echo.Context) bool {
	// Skip authentication for and signup login requests
	if e.Path() == "/favicon.ico" || e.Path() == "/api" || e.Path() == "/api/*" || e.Path() == "/api/v1/login" || e.Path() == "/api/v1/signup" {
		return true
	}
	return false
}
