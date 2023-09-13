package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"go.octolab.org/errors"
	"go.octolab.org/safe"
	"go.octolab.org/toolkit/cli/cobra"
	"go.octolab.org/unsafe"

	"go.octolab.org/template/tool/internal/command"
	"go.octolab.org/template/tool/internal/config"
)

const unknown = "unknown"

var (
	commit  = unknown
	date    = unknown
	version = "dev"
	exit    = os.Exit
	stderr  = color.Error
	stdout  = color.Output
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	root := command.New()
	root.SetErr(stderr)
	root.SetOut(stdout)
	root.AddCommand(
		cobra.NewVersionCommand(version, date, commit, config.Features...),
	)

	safe.Do(func() error { return root.ExecuteContext(ctx) }, shutdown)
}

func shutdown(err error) {
	var recovered errors.Recovered
	if errors.As(err, &recovered) {
		unsafe.DoSilent(fmt.Fprintf(stderr, "recovered: %+v\n", recovered.Cause()))
		unsafe.DoSilent(fmt.Fprintln(stderr, "---"))
		unsafe.DoSilent(fmt.Fprintf(stderr, "%+v\n", err))
	}
	exit(1)
}
