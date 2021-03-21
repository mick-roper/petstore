package root

import (
	"github.com/mick-roper/petstore/cli/cmd/pets"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "",
		Short: "Root command",
		Long:  "This is the entry point to the cli",
	}

	cmd.AddCommand(pets.NewCommand())

	return cmd
}
