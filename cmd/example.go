package exampleplugin

import (
	rootenv "github.com/ian-howell/airshipctl/pkg/environment"
	"github.com/spf13/cobra"

	"github.com/ian-howell/exampleplugin/pkg/env"
	"github.com/ian-howell/exampleplugin/pkg/example"
)

func NewExampleCommand(rootSettings *rootenv.AirshipCTLSettings) *cobra.Command {
	exampleSettings := &env.Settings{
		AirshipCTLSettings: rootSettings,
		Message: "Hello World",
	}
	exampleCommand := &cobra.Command{
		Use:   "example",
		Short: "an example plugin",
		Run: func(cmd *cobra.Command, args []string) {
			out := cmd.OutOrStdout()
			example.Fprint(out, exampleSettings.Message)
		},
	}
	return exampleCommand
}
