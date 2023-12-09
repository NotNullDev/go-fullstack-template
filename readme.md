# Go, templ, tailwindcss, daisyui, htmx, echo template

Full stack template with go (using air), templ and tailwind hot-reloading

- [Go, templ, tailwindcss, daisyui, htmx, echo template](#go-templ-tailwindcss-daisyui-htmx-echo-template)
- [Project structure](#project-structure)
  - [views](#views)
  - [notes](#notes)
  - [js-playground](#js-playground)
  - [main.go](#maingo)
  - [public](#public)
- [How to run](#how-to-run)
- [Useful documentation](#useful-documentation)
- [Fly.io](#flyio)
  - [Install flyctl](#install-flyctl)
  - [Init repo \[DONE\]](#init-repo-done)
- [How it was initialized](#how-it-was-initialized)


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

## public

Public catalog mapped to /

Files from this catalog will be publicly available.


# How to run

```shell
./devops/install-tools.sh # only for the first run
./devops/run.sh
# start coding
```


# Useful documentation

context: to avoid props drilling:
https://templ.guide/syntax-and-usage/context

using js in templates:
https://templ.guide/syntax-and-usage/script-templates

IDE support
https://templ.guide/commands-and-tools/ide-support


# Fly.io

## Install flyctl
https://fly.io/docs/hands-on/install-flyctl/

## Init repo [DONE]
https://fly.io/docs/languages-and-frameworks/golang/
```shell
fly launch ! fly.toml was copied by example because flyctl detected current project as nodejs app
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

