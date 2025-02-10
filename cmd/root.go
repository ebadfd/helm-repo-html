package cmd

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func newRootCmd(version string) *cobra.Command {

	cmd := &cobra.Command{
		Use:     "helm-repo-html",
		Short:   "Generate an HTML index for a Helm repository",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(newBuildCommand())
	return cmd
}

func Execute(version string) error {
	if err := newRootCmd(version).Execute(); err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}
	return nil
}
