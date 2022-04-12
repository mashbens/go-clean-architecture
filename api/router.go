package api

import (
	"github.com/labstack/echo/v4"

	contentV1 "clean-arch/api/v1/content"
)

type Controller struct {
	ContentV1Controller *contentV1.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	contentV1 := e.Group("/v1/content")
	contentV1.GET("", controller.ContentV1Controller.GetAll)
}
