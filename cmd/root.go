package cmd

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
)

func NewRootCmd(out io.Writer, args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "umbrella",
		Short:        "Umbrella Helm Chart utilities",
		Long:         "Umbrella Helm Chart utilities",
		SilenceUsage: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return errors.New("no arguments accepted")
			}
			return nil
		},
	}

	flags := cmd.PersistentFlags()
	flags.Parse(args)

	cmd.AddCommand(
		newLintCommand(out),
	)

	return cmd
}
