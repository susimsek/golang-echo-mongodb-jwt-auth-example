package routes

import (
	"github.com/labstack/echo/v4"
	"golang-echo-mongodb-jwt-auth-example/controller"
)

func GetUserApiRoutes(e *echo.Echo, userController *controller.UserController) {
	v1 := e.Group("/api/v1")
	{
		v1.POST("/login", userController.AuthenticateUser)
		v1.GET("/users", userController.GetAllUser)
		v1.POST("/signup", userController.SaveUser)
		v1.GET("/users/:id", userController.GetUser)
		v1.PUT("/users/:id", userController.UpdateUser)
		v1.DELETE("/users/:id", userController.DeleteUser)

	}
}
