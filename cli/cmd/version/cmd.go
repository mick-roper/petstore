package version

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	Version string
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "prints the current app version",
		Run:   run,
	}

	return cmd
}

func run(_ *cobra.Command, _ []string) {
	if Version == "" {
		log.Print("you are using an unknown version - its probably a dev version, so don't expect it to work properly :D")
		return
	}

	log.Print("Petstore CLI version ", Version)
}
