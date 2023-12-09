package main

import (
	"github.com/labstack/echo/v4"
	jsPlaygroundApi "github.com/notnulldev/go-templ-playground/js-playground/api"
	notesApi "github.com/notnulldev/go-templ-playground/notes/api"
)

func main() {
	e := echo.New()

	e.Static("/", "public")

	notesApi.InitApi(e)
	jsPlaygroundApi.InitApi(e)

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(302, "/notes")
	})

	if err := e.Start(":8080"); err != nil {
		panic(err.Error())
	}
}
