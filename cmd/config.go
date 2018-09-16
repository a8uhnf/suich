package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func ConfigCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Get current config name",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("----afas Current Config ----XXX --- goo ggoo ")
		},
	}
}
