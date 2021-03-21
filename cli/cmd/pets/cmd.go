package pets

import (
	"log"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "pets",
		Short: "pet functions",
		Long:  "pet functions",
		Run:   run,
	}

	return &cmd
}

func run(_ *cobra.Command, _ []string) {
	log.Fatal("not implemented")
}
