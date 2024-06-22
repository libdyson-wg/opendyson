package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use: "listen serial",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("must specify serial number")
		}

		iot, _ := cmd.Flags().GetBool("iot")
		return funcs.MQTTListen(args[0], iot)
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)

	listenCmd.Flags().BoolP("iot", "", false, "connect through AWS IoT instead of local MQTT")
}
