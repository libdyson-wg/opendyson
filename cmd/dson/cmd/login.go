package cmd

import (
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use: "login",
}

func init() {
	loginCmd.Run =
		rootCmd.AddCommand(loginCmd)
}
