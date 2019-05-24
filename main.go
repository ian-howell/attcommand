package main

import (
	"os"

	"github.com/ian-howell/airshipctl/cmd"
	example "github.com/ian-howell/exampleplugin/cmd"
)

func main() {
	rootCmd, settings, _ := cmd.NewRootCmd(os.Stdout)
	rootCmd.AddCommand(example.NewExampleCommand(os.Stdout, settings))
	settings.Init()
	rootCmd.Execute()
}
