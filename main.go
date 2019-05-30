package main

import (
	"os"

	"github.com/ian-howell/airshipctl/cmd"
	// "github.com/ian-howell/airshipctl/cmd/bootstrap"
	// "github.com/ian-howell/airshipctl/cmd/workflow"
	example "github.com/ian-howell/exampleplugin/cmd"
)

func main() {
	rootCmd, settings, _ := cmd.NewRootCmd(os.Stdout)
	rootCmd.AddCommand(example.NewExampleCommand(settings))

	// Optional - adds all of the default commands
	// cmd.AddDefaultAirshipCTLCommands(rootCmd, settings)

	// Optional - Pick and choose desired commands
	// rootCmd.AddCommand(bootstrap.NewBootstrapCommand(settings))
	// rootCmd.AddCommand(workflow.NewWorkflowCommand(settings))

	rootCmd.Execute()
}
