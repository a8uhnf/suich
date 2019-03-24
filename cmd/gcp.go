package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GCPConfigSwitch() *cobra.Command {
	fmt.Println("----------")
	return &cobra.Command{
		Use:   "gcp",
		Short: "Command to switch gcloud config",
		Run: func(c *cobra.Command, args []string) {

		},
	}
}
