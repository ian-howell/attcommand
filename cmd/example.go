package exampleplugin

import (
	rootenv "github.com/ian-howell/airshipctl/pkg/environment"
	"github.com/spf13/cobra"

	"github.com/ian-howell/exampleplugin/pkg/env"
	"github.com/ian-howell/exampleplugin/pkg/example"
)

const PluginSettingsID = "example"

func NewExampleCommand(rootSettings *rootenv.AirshipCTLSettings) *cobra.Command {
	exampleSettings := &env.Settings{}
	exampleCommand := &cobra.Command{
		Use:   "example",
		Short: "an example plugin",
		Run: func(cmd *cobra.Command, args []string) {
			out := cmd.OutOrStdout()
			example.Fprint(out, exampleSettings.Message)
		},
	}
	exampleSettings.Init()
	rootSettings.Register(PluginSettingsID, exampleSettings)
	return exampleCommand
}
