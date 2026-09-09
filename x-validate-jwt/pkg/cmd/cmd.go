package cmd

import (
	"os"

	"github.com/sikalabsx/go-scripts-test/x-validate-jwt/pkg/validate_jwt"
	"github.com/spf13/cobra"
)

var FlagVerbose bool

var Cmd = &cobra.Command{
	Use:   "x-validate-jwt",
	Short: "Validate JWT",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := validate_jwt.ValidateJWT(args[0])
		if err != nil {
			cmd.PrintErrln("Error validating JWT:", err)
			os.Exit(1)
		}
		if FlagVerbose {
			cmd.Println("JWT validation succeeded")
		}
	},
}

func init() {
	Cmd.PersistentFlags().BoolVarP(
		&FlagVerbose,
		"verbose",
		"v",
		false,
		"Enable verbose output",
	)
}
