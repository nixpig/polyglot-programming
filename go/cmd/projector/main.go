package main

import (
	"fmt"
	"log"

	projector "github.com/nixpig/polyglot-programming/pkg/projector"
)

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("Unable to get options %v", err)
	}

	config, err := projector.NewConfig(opts)
	if err != nil {
		log.Fatalf("Error occurred when getting new config: %v", err)
	}

	fmt.Printf("config: %+v", config)
}
