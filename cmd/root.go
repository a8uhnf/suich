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
			fmt.Println("Suich: k8s swiss knife")
		},
	}

	cmds.AddCommand(SwitchCmd())
	cmds.AddCommand(ChangeKubectl())
	cmds.AddCommand(RemoveContext())
	cmds.AddCommand(PortForward())
	cmds.AddCommand(GCPConfigSwitch())
	return cmds
}
