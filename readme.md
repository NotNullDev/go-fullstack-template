# Go, templ, tailwindcss, daisyui, htmx, echo template

Full stack template with go (using air), templ and tailwind hot-reloading

- [Go, templ, tailwindcss, daisyui, htmx, echo template](#go-templ-tailwindcss-daisyui-htmx-echo-template)
- [Project structure](#project-structure)
  - [views](#views)
  - [notes](#notes)
  - [js-playground](#js-playground)
  - [main.go](#maingo)
- [How to run](#how-to-run)
  - [First run](#first-run)
- [How it was initialized](#how-it-was-initialized)
- [Useful documentation](#useful-documentation)


# Project structure

## views

Common module, shared between other modules.

Includes global css, layout and header.

## notes

Full blown module. Example of create / read on note.

## js-playground

Really basic example of calling js from templ template.

## main.go

Entrypoint of the applicaiton. Import all other modules (notes, js-playground)

# How to run

## First run

```shell
npm i
go mod tidy
bash ./run.sh
# start coding
```

# How it was initialized

Could be useful for the future reference:

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

# Useful documentation

context: to avoid props drilling:
https://templ.guide/syntax-and-usage/context

using js in templates:
https://templ.guide/syntax-and-usage/script-templates

IDE support
https://templ.guide/commands-and-tools/ide-support
