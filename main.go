package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/notnulldev/go-templ-playground/notes"
)

var notesDb []notes.Note = []notes.Note{{Id: uuid.NewString(), Title: "Test note", Content: "Test\nMultiline\nContent"}}

func main() {
	println("hello world!")
	e := echo.New()

	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), notes.NotesContextKey, notesDb)
		notes.HomePage().Render(ctx, c.Response().Writer)
		return nil
	})

	e.GET("/create-note", func(c echo.Context) error {
		notes.CreateNotePage().Render(c.Request().Context(), c.Response().Writer)
		return nil
	})

	e.POST("/create-note", func(c echo.Context) error {
		title := c.FormValue("title")
		content := c.FormValue("content")

		notesDb = append(notesDb, notes.Note{
			Id:      uuid.NewString(),
			Title:   title,
			Content: content,
		})

		return c.Redirect(302, "/")
	})

	if err := e.Start(":8080"); err != nil {
		panic(err.Error())
	}
}
