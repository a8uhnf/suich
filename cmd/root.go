package cmd

import (
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "suich",
		Short: "Root command for switch context in k8s config",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Execute()
		},
	}

	cmds.AddCommand(SwitchCmd())
	cmds.AddCommand(ChangeKubectl())
	cmds.AddCommand(RemoveContext())
	cmds.AddCommand(PortForward())
	cmds.AddCommand(GCPConfigSwitch())
	cmds.AddCommand(SuichNamespaceCMD())
	cmds.AddCommand(GetLogsCmd())
	return cmds
}
