package cmd

import (
	"github.com/sikalabsx/go-scripts-test/another-ci-build-example/pkg/another_ci_build_example"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "another-ci-build-example",
	Run: func(cmd *cobra.Command, args []string) {
		another_ci_build_example.Another_CI_Build_Example()
	},
}
