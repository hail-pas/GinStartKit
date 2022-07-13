SHELL = /bin/bash

CONFIG_FILE         = config/content/default.yaml
BUILD_IMAGE  = golang:1.16
IMAGE_NAME          = GinStartKit
REPOSITORY          = rep
build:
	go build -tags=jsoniter .