package cmd

import (
	"fmt"

	"github.com/ebadfd/helm-repo-html/pkg/builder"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

var (
	BuildName  = "build"
	BuildShort = "Generate an HTML index for a Helm repository"
)

type buildCommand struct {
	Template string
	Input    string `validate:"required"`
	Output   string
}

func defaultBuildCommandOptions() *buildCommand {
	return &buildCommand{
		Output: "-",
	}
}

func newBuildCommand() *cobra.Command {
	c := defaultBuildCommandOptions()

	cmd := &cobra.Command{
		Use:   c.Name(),
		Short: c.Short(),
		RunE:  c.run,
	}

	cmd.Flags().StringVarP(&c.Template, "template", "t", c.Template, "template file")
	cmd.Flags().StringVarP(&c.Input, "input", "i", c.Input, "input file name")
	cmd.Flags().StringVarP(&c.Output, "output", "o", c.Output, "output file name")

	return cmd
}

func (o *buildCommand) run(cmd *cobra.Command, args []string) error {
	err := validate.Struct(o)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		fmt.Println("Validation errors:")
		for _, ve := range err.(validator.ValidationErrors) {
			fmt.Printf("  - Field: %s\n", ve.Field())
			fmt.Printf("    Issue: Validation '%s' failed\n", ve.Tag())
			if ve.Param() != "" {
				fmt.Printf("    Expected: %s\n", ve.Param())
			}
			fmt.Printf("    Actual: %v\n\n", ve.Value())
		}

		return fmt.Errorf("validation failed: %d errors found", len(err.(validator.ValidationErrors)))
	}

	err = builder.Handler(o.Template, &o.Input, o.Output)
	if err != nil {
		panic(err)
	}
	return nil
}

func (c *buildCommand) Name() string {
	return BuildName
}

func (c *buildCommand) Short() string {
	return BuildShort
}
