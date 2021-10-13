package command

import (
	"github.com/spf13/cobra"

	"go.octolab.org/template/tool/internal/command/demo"
)

// New returns the new root command.
func New() *cobra.Command {
	command := cobra.Command{
		Use:   "%template%",
		Short: "%template%",
		Long:  "%template%",

		Args: cobra.NoArgs,

		SilenceErrors: false,
		SilenceUsage:  true,
	}

	/* configure instance */
	command.AddCommand(
		demo.Panic(),
		demo.Stderr(),
		demo.Stdout(),
	)

	return &command
}
