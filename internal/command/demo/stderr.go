package demo

import (
	"strings"

	"github.com/spf13/cobra"
)

// Stderr returns a demo cobra.Command to write to the stderr.
//
//	$ go run main.go stderr [message]
//	$ go run main.go stderr [message] >/dev/null
//	$ go run main.go stderr [message] 2>/dev/null
func Stderr() *cobra.Command {
	command := cobra.Command{
		Use: "stderr",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				args = []string{cmd.Name(), "content"}
			}
			cmd.PrintErrln(strings.Join(args, " "))
		},
	}

	return &command
}
