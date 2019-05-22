package cmd

import (
	"github.com/spf13/cobra"
)

func GCPConfigSwitch() *cobra.Command {
	return &cobra.Command{
		Use:   "gcp",
		Short: "Command to switch gcloud config[IN PROGRESS]",
		Run: func(c *cobra.Command, args []string) {

		},
	}
}
