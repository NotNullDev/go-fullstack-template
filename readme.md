# How to run

```shell
npm i
go mod tidy
bash ./run.sh
# start coding
```

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
context: to avoid props drilling:
https://templ.guide/syntax-and-usage/context

using js in templates:
https://templ.guide/syntax-and-usage/script-templates

https://templ.guide/commands-and-tools/ide-support