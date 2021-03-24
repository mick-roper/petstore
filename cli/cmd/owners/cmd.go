package owners

import (
	"log"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "owners",
		Short: "Modify owner info",
		Run:   run,
	}

	return cmd
}

func run(_ *cobra.Command, _ []string) {
	log.Fatal("not implemented")
}
