package api

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/notnulldev/go-templ-playground/notes/types"
	"github.com/notnulldev/go-templ-playground/notes/views"
)

var notesDb []types.Note = []types.Note{{Id: uuid.NewString(), Title: "Test note", Content: "Test\nMultiline\nContent"}}

func InitApi(e *echo.Echo) {
	e.GET("/notes", func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), types.NotesContextKey, notesDb)
		views.HomePage().Render(ctx, c.Response().Writer)
		return nil
	})

	e.GET("/notes/create", func(c echo.Context) error {
		views.CreateNotePage().Render(c.Request().Context(), c.Response().Writer)
		return nil
	})

	e.POST("/notes/create", func(c echo.Context) error {
		title := c.FormValue("title")
		content := c.FormValue("content")

		notesDb = append(notesDb, types.Note{
			Id:      uuid.NewString(),
			Title:   title,
			Content: content,
		})

		return c.Redirect(302, "/")
	})
}
