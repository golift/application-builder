package main

import (
	"log"

	// This path doesn't match the repo name. Use the correct name locally.
	"github.com/golift/hello-world/helloworld"
)

// Keep it simple.
func main() {
	if err := helloworld.Start(); err != nil {
		log.Fatalln("[ERROR]", err)
	}
}
