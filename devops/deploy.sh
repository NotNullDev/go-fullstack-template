#!/bin/bash

# script should be run from the root directory!: `./devops/deploy.sh`

# TODO: tailwind to prod mode
go build -o go-fullstack .
docker build -t go-fullstack:latest .
fly deploy