package cmd

import (
	"fmt"

	"gopkg.in/yaml.v3"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		ds, err := funcs.GetDevices()
		if err != nil {
			return err
		}

		fmt.Println("Available devices:")
		bs, err := yaml.Marshal(ds)
		fmt.Println(string(bs))
		return err
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println("This is sensitive information that can be used to control your devices. " +
			"Please take caution before sharing the information above with anyone you do not trust.")
	},
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}
