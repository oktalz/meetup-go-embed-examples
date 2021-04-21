package main

import (
	"embed"
	"fmt"
)

// START OMIT

//go:embed go/*
var fs embed.FS

func main() {

	fmt.Println("-- go.mod --")
	data, _ := fs.ReadFile("go/go.mod")
	fmt.Println(string(data))

	fmt.Println("-- hello.go --")
	data, _ = fs.ReadFile("go/hello.go")
	fmt.Println(string(data))
}

// END OMIT
