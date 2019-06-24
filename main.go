package main

import (
	"os"

	"github.com/ian-howell/airshipctl/cmd"
	// "github.com/ian-howell/airshipctl/cmd/bootstrap"
	example "github.com/ian-howell/exampleplugin/cmd"
)

func main() {
	// Choose one of these
	// 1. Creates an airshipctl with all of the default commands
	rootCmd, settings, _ := cmd.NewAirshipCTLCommand(os.Stdout)
	// 2. Creates a bare-bones variant of airshipctl, allowing the user to
	// pick and choose which subcommands he wants
	// rootCmd, settings, _ := cmd.NewRootCmd(os.Stdout)

	// This line adds the "plugin"
	rootCmd.AddCommand(example.NewExampleCommand(settings))

	// Optional - Pick and choose desired commands - Only necessary if
	// option 2 was used for root command creation
	// rootCmd.AddCommand(bootstrap.NewBootstrapCommand(settings))

	rootCmd.Execute()
}
