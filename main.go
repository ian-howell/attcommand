package main

import (
	"os"

	"github.com/ian-howell/airshipctl/cmd"
	example "github.com/ian-howell/exampleplugin/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(os.Stdout)
	rootCmd.AddCommand(example.NewExampleCommand(os.Stdout, os.Args[1:]))
	rootCmd.Execute()
}
