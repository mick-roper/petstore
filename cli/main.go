package main

import (
	"log"

	root "github.com/mick-roper/petstore/cli/cmd"
)

var (
	Version string
)

func main() {
	log.Fatal(root.NewRootCommand().Execute())
}
