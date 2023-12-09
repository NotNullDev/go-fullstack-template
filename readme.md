```shell
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/cosmtrek/air@latest
go get github.com/labstack/echo/v4
```

```shell
air init

npm init -y
npm install -D tailwindcss
npx tailwindcss init

templ generate -watch
npx tailwindcss -i ./views/app.css -o ./dist/output.css --watch
air
```

https://templ.guide/commands-and-tools/ide-support