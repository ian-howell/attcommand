package exampleplugin

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/ian-howell/exampleplugin/pkg/example"
)

func NewExampleCommand(out io.Writer, args []string) *cobra.Command {
	exampleCommand := &cobra.Command{
		Use:   "example",
		Short: "an example plugin",
		Run: func(cmd *cobra.Command, args []string) {
			example.Print()
		},
	}
	return exampleCommand
}
