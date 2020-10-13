package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/types"

	"github.com/a8uhnf/suich/pkg/utils"
)

func patch() *cobra.Command {
	var deployment string
	var namespace string
	c := &cobra.Command{
		Use:   "patch",
		Short: "patch deployment with date or any other cause",
		RunE: func(cmd *cobra.Command, args []string) error {
			kc := utils.GetClientOutOfCluster()
			date := time.Now().Format("2006-01-02T15:04:05.999999-07:00")
			p := fmt.Sprintf("{\"spec\":{\"template\":{\"metadata\":{\"annotations\":{\"date\":\"%s\"}}}}}", date)

			d, err := kc.AppsV1().Deployments(namespace).Patch(deployment, types.StrategicMergePatchType, []byte(p))
			if err != nil {
				return err
			}
			fmt.Printf("%s patched", d.ObjectMeta.Name)
			return nil
		},
	}
	c.Flags().StringVarP(&deployment, "deploy", "d", "", "deployment name")
	c.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace name")
	return c
}
