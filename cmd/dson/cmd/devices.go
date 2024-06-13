package cmd

import (
	"github.com/libdyson-wg/libdyson-go/config"

	"github.com/spf13/cobra"
)

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Lists the devices on your account",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		tok, err := config.GetToken()
		if err != nil {
			return err
		}
		if tok == "" {
			err = funcs.Login()
		}
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		//ds, err := funcs.GetDevices
	},
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}
