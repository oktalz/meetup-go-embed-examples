package main

import (
	"embed"
	"fmt"
)

// START OMIT

//go:embed c python go
var fs embed.FS

func main() {
	fmt.Println("-- c --")
	data, _ := fs.ReadFile("c/hello.c")
	fmt.Println(string(data))

	fmt.Println("-- py --")
	data, _ = fs.ReadFile("python/hello.py")
	fmt.Println(string(data))

	fmt.Println("-- go --")
	data, _ = fs.ReadFile("go/hello.go")
	fmt.Println(string(data))
}

// END OMIT
