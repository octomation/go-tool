package main

import (
	"context"
	"os"

	"go.octolab.org/errors"
	"go.octolab.org/safe"
	"go.octolab.org/toolkit/cli/cobra"

	"tool/internal/cmd"
)

const unknown = "unknown"

var (
	commit  = unknown
	date    = unknown
	version = "dev"
)

func main() {
	ctx := context.TODO()
	root := cmd.New()
	root.SetOut(os.Stdout)
	root.SetErr(os.Stderr)
	root.AddCommand(
		cobra.NewCompletionCommand(),
		cobra.NewVersionCommand(version, date, commit),
	)
	safe.Do(func() error {
		return root.ExecuteContext(ctx)
	}, func(err error) {
		if recovered, is := errors.Unwrap(err).(errors.Recovered); is {
			root.PrintErrf("recovered: %+v\n", recovered.Cause())
			root.PrintErrln("---")
			root.PrintErrf("%+v\n", err)
		}
		os.Exit(1)
	})
}
