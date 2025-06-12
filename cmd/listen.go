 package cmd
 
 import (
 	"fmt"
 	"github.com/spf13/cobra"
 )
 
 // listenCmd represents the listen command
 var listenCmd = &cobra.Command{
-	Use:   "listen serial",
+	Use:   "listen serial|ALL",
 	Short: "Continuously listen to and print messages for a device",
-	Long: "Continuously listen to and print messages for a device. This uses MQTT over LAN by default. But if needed, " +
-		"you can use the -iot flag to use the official Dyson cloud setup instead (AWS IoT).",
+	Long: "Continuously listen to and print messages for a device. This uses MQTT over LAN by default. " +
+		"Specify `ALL` as the serial to subscribe to every discovered device. " +
+		"If needed, you can use the -iot flag to use the official Dyson cloud setup instead (AWS IoT).",
 	RunE: func(cmd *cobra.Command, args []string) error {
 		if len(args) < 1 {
 			return fmt.Errorf("must specify serial")
 		}
 
 		iot, _ := cmd.Flags().GetBool("iot")
 		return funcs.MQTTListen(args[0], iot)
 	},
 }
 
 func init() {
 	rootCmd.AddCommand(listenCmd)
 
 	listenCmd.Flags().BoolP("iot", "", false, "connect through AWS IoT instead of local MQTT")
 }
