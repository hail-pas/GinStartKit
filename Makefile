SHELL = /bin/bash

CONFIG_FILE         = config/content/default.yaml
BUILD_IMAGE         = golang:1.18.7
IMAGE_NAME          = GinStartKit
REPOSITORY          = rep

migrations:
	migrate create -ext sql -dir storage/relational/migrations -seq $(label)
	@echo "Success"
build:
	go build -race -tags=jsoniter -o server core/server.go

clear:
	# zsh
	# setopt extendedglob && rm -rf ^server
	# bash
	# bash -O extglob
	shopt -s extglob && rm -rf !(server)