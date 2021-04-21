package main

import (
	"embed"
	"fmt"
)

//go:embed test
var fs embed.FS

func main() {
	data, _ := fs.ReadFile("test/test.txt")
	fmt.Print(string(data))
}
