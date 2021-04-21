all: build

.PHONY: build
build:
	-mv go/go.mod go/go.mod.goembed
	go build -o example
	-mv go/go.mod.goembed go/go.mod
