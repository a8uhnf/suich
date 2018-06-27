package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func SwitchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "switch",
		Short: "To switch context use this command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Switch!!")
		},
	}
}
