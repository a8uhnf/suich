package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "suich",
		Short: "Root command for switch context in k8s config",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello World!!!")
		},
	}

	cmds.AddCommand(SwitchCmd())
	return cmds
}
