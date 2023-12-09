package main

import (
	"github.com/labstack/echo/v4"
	"github.com/notnulldev/go-templ-playground/views"
)

func main() {
	println("hello world!")
	e := echo.New()

	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "hello!")
	})

	e.GET("/component", func(c echo.Context) error {
		views.HomePage().Render(c.Request().Context(), c.Response().Writer)
		return nil
	})

	if err := e.Start(":8080"); err != nil {
		panic(err.Error())
	}
}
