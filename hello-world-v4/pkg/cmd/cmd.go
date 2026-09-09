package cmd

import (
	"github.com/sikalabsx/go-scripts-test/hello-world-v2/pkg/hello_world"
	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:   "hello-world-v4",
	Short: "Hello World example",
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
