package content

import (
	"clean-arch/api/v1/content/request"
	contentBusiness "clean-arch/business/content"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service contentBusiness.Service
}

func NewController(service contentBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {

	contents, err := controller.service.GetContents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, contents)
}

func (controller *Controller) GetContentByID(c echo.Context) error {
	req := new(request.Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	content, err := controller.service.GetContentByID(req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, content)
}
