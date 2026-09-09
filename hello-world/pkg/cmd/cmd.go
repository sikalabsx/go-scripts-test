package cmd

import (
	"github.com/sikalabsx/go-scripts-test/hello-world/pkg/hello_world"
	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:   "hello-world",
	Short: "Hello World Example",
	Run: func(cmd *cobra.Command, args []string) {
		hello_world.PrintHelloWorld(FlagName)
	},
}

func init() {
	Cmd.Flags().StringVarP(
		&FlagName,
		"name",
		"n",
		"World",
		"Name to greet",
	)
}
