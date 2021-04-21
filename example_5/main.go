package main

import (
	"embed"
	"fmt"
)

// START OMIT

//go:embed test assets
var fs embed.FS

func main() {
	data, _ := fs.ReadFile("test/test.txt")
	fmt.Print(string(data))
	data, _ = fs.ReadFile("assets/assets.txt")
	fmt.Print(string(data))
}

// END OMIT
