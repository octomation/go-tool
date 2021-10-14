package demo

import (
	"strings"

	"github.com/spf13/cobra"
)

// Panic returns a demo cobra.Command to raise a panic.
//
//	$ go run main.go panic [message]
func Panic() *cobra.Command {
	command := cobra.Command{
		Use: "panic",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				args = []string{"unexpected", cmd.Name()}
			}
			panic(strings.Join(args, " "))
		},
	}

	return &command
}
