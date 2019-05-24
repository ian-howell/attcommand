package exampleplugin

import (
	"io"

	"github.com/spf13/cobra"
	rootenv "github.com/ian-howell/airshipctl/pkg/environment"

	"github.com/ian-howell/exampleplugin/pkg/example"
	"github.com/ian-howell/exampleplugin/pkg/env"
)

const PluginSettingsID = "example"

func NewExampleCommand(out io.Writer, rootSettings *rootenv.AirshipCTLSettings) *cobra.Command {
	exampleSettings := &env.Settings{}
	exampleCommand := &cobra.Command{
		Use:   "example",
		Short: "an example plugin",
		Run: func(cmd *cobra.Command, args []string) {
			example.Print(exampleSettings.Message)
		},
	}
	exampleSettings.Init()
	rootSettings.Register(PluginSettingsID, exampleSettings)
	return exampleCommand
}
