package demo

import "github.com/spf13/cobra"

// Panic returns a demo cobra.Command to raise a panic.
//
// 	$ go run main.go panic
//
func Panic() *cobra.Command {
	command := cobra.Command{
		Use:  "panic",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			panic("unexpected panic")
		},
	}

	return &command
}
