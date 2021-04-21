package main

import (
	"log"
	"net/http"
)

// START OMIT

func main() {
	fs := http.FileServer(http.Dir("./ui/dist"))
	http.Handle("/", fs)

	log.Println("Listen on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// END OMIT
