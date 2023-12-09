#!/bin/bash

templ generate -watch &
npx tailwindcss -i ./views/app.css -o ./public/app.css --watch &
air &