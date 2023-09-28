package main

import (
	"fmt"
	"log"

	projector "github.com/nixpig/polyglot-programming/pkg/config"
)

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("Unable to get options %v", err)
	}

	fmt.Printf("opts: %+v", opts)
}
