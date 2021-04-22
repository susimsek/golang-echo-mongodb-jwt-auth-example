package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang-echo-mongodb-jwt-auth-example/config"
	"golang-echo-mongodb-jwt-auth-example/controller"
	_ "golang-echo-mongodb-jwt-auth-example/docs"
	"golang-echo-mongodb-jwt-auth-example/handler"
	"golang-echo-mongodb-jwt-auth-example/repository"
	"golang-echo-mongodb-jwt-auth-example/routes"
	"golang-echo-mongodb-jwt-auth-example/util"
	"log"
)

// @title Golang User REST API
// @description Provides access to the core features of Golang User REST API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	mongoConnection, errorMongoConn := config.MongoConnection()

	if errorMongoConn != nil {
		log.Println("Error when connect mongo : ", errorMongoConn.Error())
	}

	e := echo.New()

	e.HTTPErrorHandler = handler.ErrorHandler
	e.Validator = util.NewValidationUtil()
	config.CORSConfig(e)
	config.WebSecurityConfig(e)

	userRepo := repository.NewUserRepository(mongoConnection)
	userController := controller.NewUserController(userRepo)
	routes.GetUserApiRoutes(e, userController)
	routes.GetSwaggerRoutes(e)

	// echo server 9000 de başlatıldı.
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ServerPort)))
}
