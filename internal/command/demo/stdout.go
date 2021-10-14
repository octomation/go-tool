package demo

import (
	"strings"

	"github.com/spf13/cobra"
)

// Stdout returns a demo cobra.Command to write to the stdout.
//
//	$ go run main.go stdout [message]
//	$ go run main.go stdout [message] >/dev/null
func Stdout() *cobra.Command {
	command := cobra.Command{
		Use: "stdout",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				args = []string{cmd.Name(), "content"}
			}
			cmd.Println(strings.Join(args, " "))
		},
	}

	return &command
}
