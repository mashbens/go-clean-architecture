package api

import (
	"github.com/labstack/echo/v4"

	"clean-arch/api/middleware"
	auth "clean-arch/api/v1/auth"
	contentV1 "clean-arch/api/v1/content"
)

type Controller struct {
	ContentV1Controller *contentV1.Controller

	AuthController *auth.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	contentV1 := e.Group("/v1/content")
	contentV1.Use(middleware.JWTMiddleware())

	contentV1.GET("", controller.ContentV1Controller.GetAll)
	contentV1.GET("/:id", controller.ContentV1Controller.GetContentByID)
	contentV1.POST("", controller.ContentV1Controller.CreateContent)
	contentV1.PUT("/:id", controller.ContentV1Controller.UpdateContent)

	auth := e.Group("/v1/auth")
	auth.POST("", controller.AuthController.Auth)
}
