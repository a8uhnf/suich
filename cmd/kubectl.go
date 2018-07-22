package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ChangeKubectl() *cobra.Command {
	var version string
	cmd := &cobra.Command{
		Use:   "kubectl",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			//cmd.Flags().
			fmt.Println("Hello World!!!")
			k := &Kubectl{}
			err := k.Validate(cmd)
			if err != nil {
				panic(err)
			}
		},
	}
	cmd.Flags().StringVarP(&version, "version", "v", "", "kubectl valid version")
	return cmd
}

type Kubectl struct{}

func (k *Kubectl) Validate(cmd *cobra.Command) error {
	v := cmd.Flag("version").Value
	if v.String() == "" {
		return fmt.Errorf("%s name must provided", "version")
	}
	return nil
}
