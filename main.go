package main

import (
	"log"

	"github.com/golift/application-builder/helloworld"
)

// Keep it simple.
func main() {
	if err := helloworld.Start(); err != nil {
		log.Fatalln("[ERROR]", err)
	}
}
