package main

import (
	"log"

	root "github.com/mick-roper/petstore/cli/cmd"
)

func main() {
	log.Fatal(root.NewRootCommand().Execute())
}
