package main

import (
	"embed"
	"fmt"
)

// START OMIT

//go:embed assets/*
var fs embed.FS

func main() {
	fmt.Println("Hello world !")
}

// END OMIT
