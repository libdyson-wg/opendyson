package cmd

import (
	"fmt"

	"github.com/libdyson-wg/opendyson/internal/config"

	"gopkg.in/yaml.v3"

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
		fmt.Println("This is sensitive information that can be used to remotely control your Dyson devices. " +
			"Please take caution before sharing it with anyone you do not trust.")
	},
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}
