package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/a8uhnf/suich/pkg/api"
	meta_v1"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func SuichNamespaceCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:"ns",
		Run: func(cmd *cobra.Command, args []string) {
			fixNamespace()
		},
		Short:"fix specific namespace",
	}

	cmd.Flags().StringP("namespace", "n", "default", "set namespace to specified namespace.")
	return cmd
}

func fixNamespace()  {
	c := api.GetKubernetesClient()
	list, err := c.CoreV1().Namespaces().List(meta_v1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
		return
	}
	namespaces := api.GetNamespaceNames(list)

	str , err := api.RunPrompt(namespaces, "select namespace")
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
	err = modifyContext(str)
	if err != nil {
		panic(err)
	}
}


func modifyContext(ns string) error {
	pathOptions := clientcmd.NewDefaultPathOptions()
	cfg, err := pathOptions.GetStartingConfig()
	if err != nil {
		return err
	}
	name := cfg.CurrentContext
	ctx := cfg.Contexts[name]
	fmt.Println("current context")
	fmt.Println(ctx)
	ctx.Namespace = ns
	cfg.Contexts[name] = ctx
	err = clientcmd.ModifyConfig(pathOptions, *cfg, true)
	if err != nil {
		return err
	}
	return nil
}