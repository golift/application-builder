package main

import (
	"log"

	// This path doesn't match the repo name. Use the correct name locally.
	"golift.io/application-builder/helloworld"
)

// Keep it simple.
func main() {
	if err := helloworld.Start(); err != nil {
		log.Fatalln("[ERROR]", err)
	}
}
