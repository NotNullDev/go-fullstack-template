package api

import (
	"github.com/labstack/echo/v4"
	"github.com/notnulldev/go-templ-playground/js-playground/views"
)

func InitApi(e *echo.Echo) {
	e.GET("/js-playground", func(c echo.Context) error {
		views.CreateJSPlaygroundPage().Render(c.Request().Context(), c.Response().Writer)
		return nil
	})
}
