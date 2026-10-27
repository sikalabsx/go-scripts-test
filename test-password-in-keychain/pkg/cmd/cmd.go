package cmd

import (
	"github.com/sikalabsx/go-scripts-test/test-password-in-keychain/pkg/test_password_in_keychain"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "test-password-in-keychain",
	Short: "Set and get a test password in the system keychain",
}

var SetCmd = &cobra.Command{
	Use:   "set <password>",
	Short: "Store the password in the keychain",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		test_password_in_keychain.SetPassword(args[0])
	},
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Print the password from the keychain",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		test_password_in_keychain.PrintPassword()
	},
}

func init() {
	Cmd.AddCommand(SetCmd)
	Cmd.AddCommand(GetCmd)
}
