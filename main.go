package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	jsPlaygroundApi "github.com/notnulldev/go-templ-playground/js-playground/api"
	notesApi "github.com/notnulldev/go-templ-playground/notes/api"
)

//go:embed public/*
var publicFiles embed.FS

func main() {
	e := echo.New()

	initFs(e)

	notesApi.InitApi(e)
	jsPlaygroundApi.InitApi(e)

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(302, "/notes")
	})

	if err := e.Start(":8080"); err != nil {
		panic(err.Error())
	}
}

func initFs(e *echo.Echo) {
	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	assetHandler := http.FileServer(getFileSystem(useOS))
	e.GET("/", echo.WrapHandler(assetHandler))
}

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("app"))
	}

	log.Print("using embed mode")
	fsys, err := fs.Sub(publicFiles, "app")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
