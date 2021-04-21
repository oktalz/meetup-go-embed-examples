package main

import (
	"embed"
	"fmt"
	"log"
	"path"
	"strings"
)

// START OMIT

//go:embed go/*
var fs embed.FS

func newName(name string) string {
	if strings.HasSuffix(name, ".goembed") {
		return strings.TrimSuffix(name, ".goembed")
	}
	return name
}

func main() {

	dirEntries, err := fs.ReadDir("go")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}
		fmt.Println(path.Join("go", newName(entry.Name())))
	}
}

// END OMIT
