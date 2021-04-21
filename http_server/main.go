package main

import (
	"embed"
	"log"
	"net/http"
)

// START OMIT

//go:embed ui/dist/*
var staticFiles embed.FS

func main() {
	var staticFS = http.FS(staticFiles)
	fs := http.FileServer(staticFS)

	// Serve static files
	http.Handle("/", fs)

	log.Println("Listen on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// END OMIT
