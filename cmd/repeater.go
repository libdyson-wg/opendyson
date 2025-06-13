package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var repeaterCmd = &cobra.Command{
	Use:   "repeater serial|ALL",
	Short: "Repeat device MQTT messages to a remote broker",
	Long:  "Subscribe to device messages and publish them to a remote MQTT broker.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("must specify serial")
		}
		iot, _ := cmd.Flags().GetBool("iot")
		host, _ := cmd.Flags().GetString("host")
		if host == "" {
			return fmt.Errorf("host must be specified")
		}
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		return funcs.MQTTRepeater(args[0], iot, host, user, password)
	},
}

func init() {
	rootCmd.AddCommand(repeaterCmd)
	repeaterCmd.Flags().BoolP("iot", "", false, "connect through AWS IoT instead of local MQTT")
	repeaterCmd.Flags().StringP("host", "", "", "remote MQTT host (hostname or address)")
	repeaterCmd.Flags().StringP("user", "u", "", "username for remote MQTT host")
	repeaterCmd.Flags().StringP("password", "p", "", "password for remote MQTT host")
}
