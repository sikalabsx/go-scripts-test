package cmd

import (
	"github.com/sikalabsx/go-scripts-test/gravatar-image-from-email/pkg/gravatar_image_from_email"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "gravatar-image-from-email <email>",
	Short: "Print the Gravatar avatar URL for an email address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gravatar_image_from_email.PrintGravatarImageURL(args[0])
	},
}
