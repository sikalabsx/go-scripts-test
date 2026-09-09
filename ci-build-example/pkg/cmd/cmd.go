package cmd

import (
	"github.com/sikalabsx/go-scripts-test/ci-build-example/pkg/ci_build_example"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "ci-build-example",
	Run: func(cmd *cobra.Command, args []string) {
		ci_build_example.CI_Build_Example()
	},
}
