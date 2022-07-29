package assets

import "embed"

var (
	//go:embed files/*
	Files embed.FS

	//go:embed templates/*
	Templates embed.FS
)
