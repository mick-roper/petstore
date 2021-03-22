package main

import (
	"log"

	"github.com/mick-roper/petstore/cli/cmd/pets"
	"github.com/mick-roper/petstore/cli/cmd/version"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

func main() {
	rootCmd := cobra.Command{}
	rootCmd.AddCommand(version.NewCommand())
	rootCmd.AddCommand(pets.NewCommand())
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
