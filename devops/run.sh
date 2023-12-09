#!/bin/bash

# script should be run from the root directory!: `./devops/run.sh`

./devops/clear.sh

templ generate -watch &
npx tailwindcss -i ./views/app.css -o ./public/app.css --watch &
air

echo "tools are running in the background"