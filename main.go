package main

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/bpdunni/helm-umbrella/cmd"
	"github.com/spf13/cobra"
)

var version string

func main() {
	umbrellaCmd := cmd.NewRootCmd(os.Stdout, os.Args[1:])
	umbrellaCmd.AddCommand(newVersionCommand(os.Stdout))

	if err := umbrellaCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func newVersionCommand(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the version number of helm umbrella",
		Long:  "Prints the version number of helm umbrella",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(out, "helm umbrella plugin version", version, "built with go version", runtime.Version())
		},
	}

	return cmd
}
