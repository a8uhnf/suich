package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)


func RootCmd() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "suich",
		Short: "Root command for switch context in k8s config",
		Long:  "",
		PreRun: func(cmd *cobra.Command, args []string) {
			ok, err := cmd.Flags().GetBool("debug")
			fmt.Println(ok, err)
			if ok, err := cmd.Flags().GetBool("debug"); err == nil &&ok {
				fmt.Println("-------------")
				log.SetOutput(os.Stdout)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("------")
			// cmd.Execute()
			cmd.Help()
		},
	}
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	cmds.Flags().Bool("debug", false, "turn on debug mode")
	cmds.AddCommand(SwitchCmd())
	cmds.AddCommand(ChangeKubectl())
	cmds.AddCommand(RemoveContext())
	cmds.AddCommand(PortForward())
	cmds.AddCommand(GCPConfigSwitch())
	cmds.AddCommand(SuichNamespaceCMD())
	cmds.AddCommand(GetLogsCmd())
	return cmds
}
