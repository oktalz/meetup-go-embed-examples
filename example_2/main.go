package main

import (
	_ "embed"
	"fmt"
)

//go:embed test/test.txt
var test []byte

func main() {
	fmt.Print(string(test))
}
