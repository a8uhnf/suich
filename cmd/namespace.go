package cmd

import (
	"github.com/spf13/cobra"
)

func SuichNamespaceCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:"ns",
		Run: func(cmd *cobra.Command, args []string) {

			// fmt.Println(f.Parse())
		},
		Short:"fixed specific namespace",
	}

	cmd.Flags().StringP("namespace", "n", "default", "set namespace to specified namespace.")
	return cmd
}

func fixNamespace()  {

}
