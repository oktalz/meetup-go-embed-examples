package main

import (
	"embed"
	"log"
	"net/http"
)

// START OMIT

//go:embed ui/dist/*
var staticFiles embed.FS

//go:embed /ui/dist/index.html
var indexPage []byte

func main() {
	var staticFS = http.FS(staticFiles)
	fs := http.FileServer(staticFS)
	...
	r.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := staticFiles.ReadFile("ui/dist" + r.RequestURI)
		if err != nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = w.Write(indexPage) // ignore since its example
			return
		}
		w.Header().Set("Content-Type", ct) // ct is calculated content type not shown here
		_, _ = w.Write(file)               // ignore since its example
	}))
	...
}

// END OMIT
