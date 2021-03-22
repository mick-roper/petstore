package main

import (
	"log"

	"github.com/mick-roper/petstore/cli/cmd/pets"
	"github.com/mick-roper/petstore/cli/cmd/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(version.NewCommand())
	rootCmd.AddCommand(pets.NewCommand())

	cobra.OnInitialize()
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
