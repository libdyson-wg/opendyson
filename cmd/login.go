package cmd

import (
	"fmt"
	"github.com/libdyson-wg/opendyson/internal/config"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your Dyson account",
	RunE: func(cmd *cobra.Command, args []string) error {
		return funcs.Login()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			fmt.Sprintf(
				"You are logged in. Please note that your Dyson API Token has been saved to %s.\n\n"+
					"This API Token is sensitive and should never be shared with anyone you don't trust. "+
					"It could possibly be used to control your Dyson devices remotely, or learn sensitive "+
					"private information about you through your Dyson account.",
				config.GetFilePath(),
			),
		)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
