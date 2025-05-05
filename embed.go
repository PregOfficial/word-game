package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed word-game-web/dist/*
var frontend embed.FS

func ServeFrontend() http.Handler {
	content, err := fs.Sub(frontend, "word-game-web/dist")
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(content))
}
