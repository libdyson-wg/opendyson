package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host serial|ALL",
	Short: "Host an MQTT server relaying device messages",
	Long:  "Start a local MQTT server on port 1883 and publish device messages to it.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("must specify serial")
		}
		iot, _ := cmd.Flags().GetBool("iot")
		return funcs.MQTTHost(args[0], iot)
	},
}

func init() {
	rootCmd.AddCommand(hostCmd)
	hostCmd.Flags().BoolP("iot", "", false, "connect through AWS IoT instead of local MQTT")
}
