package cmd

import (
	"github.com/spf13/cobra"
)

var addons_configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configures the addon w/ADDON_NAME within minikube (example: minikube addons configure registry-creds). For a list of available addons use: minikube addons list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	addonsCmd.AddCommand(addons_configureCmd)
}
