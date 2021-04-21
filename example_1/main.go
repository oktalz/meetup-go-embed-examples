package main

import (
	_ "embed"
	"fmt"
)

//go:embed test/test.txt
var test string

func main() {
	fmt.Print(test)
}
