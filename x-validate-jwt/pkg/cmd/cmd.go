package cmd

import (
	"io"
	"os"
	"strings"

	"github.com/sikalabsx/go-scripts-test/x-validate-jwt/pkg/validate_jwt"
	"github.com/spf13/cobra"
)

var FlagVerbose bool

var Cmd = &cobra.Command{
	Use:   "x-validate-jwt",
	Short: "Validate JWT",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		if token == "-" {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				cmd.PrintErrln("Error reading token from stdin:", err)
				os.Exit(1)
			}
			token = strings.TrimSpace(string(data))
		}
		err := validate_jwt.ValidateJWT(token)
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
