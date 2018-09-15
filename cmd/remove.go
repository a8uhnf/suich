package cmd

import (

	// k8s "k8s.io/kubernetes"
	"bytes"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"k8s.io/client-go/tools/clientcmd"
)

func RemoveContext() *cobra.Command {
	return &cobra.Command{
		Use:   "rm",
		Short: "Remove context and cluster from kubeconfig",
		Run: func(cmd *cobra.Command, args []string) {
			ctxs, err := readKubeConfigFile()
			if err != nil {
				log.Fatalln(err)
			}

			selectedCtx, err := runPrompt(ctxs)
			if err != nil {
				log.Fatalln(err)
			}
			if selectedCtx == "" {
				return
			}

			err = deleteClusterName(cmd, selectedCtx)
			if err != nil {
				panic(err)
			}

			err = deleteContext(cmd, selectedCtx)
			if err != nil {
				panic(err)
			}

			fmt.Println("----- ", selectedCtx)
		},
	}
}

func deleteClusterName(cmd *cobra.Command, name string) error {
	out := bytes.NewBuffer([]byte{})
	// errOut := bytes.NewBuffer([]byte{})
	configAccess := clientcmd.NewDefaultPathOptions()
	config, err := configAccess.GetStartingConfig()
	if err != nil {
		return err
	}

	args := cmd.Flags().Args()
	if len(args) == 1 {
		cmd.Help()
		return nil
	}

	configFile := configAccess.GetDefaultFilename()
	if configAccess.IsExplicitFile() {
		configFile = configAccess.GetExplicitFile()
	}

	_, ok := config.Clusters[name]
	if !ok {
		return fmt.Errorf("cannot delete cluster %s, not in %s", name, configFile)
	}

	delete(config.Clusters, name)

	if err := clientcmd.ModifyConfig(configAccess, *config, true); err != nil {
		return err
	}

	fmt.Fprintf(out, "deleted cluster %s from %s\n", name, configFile)

	return nil
}

func deleteContext(cmd *cobra.Command, name string) error {
	out := bytes.NewBuffer([]byte{})
	errOut := bytes.NewBuffer([]byte{})
	configAccess := clientcmd.NewDefaultPathOptions()
	config, err := configAccess.GetStartingConfig()
	if err != nil {
		return err
	}

	args := cmd.Flags().Args()
	if len(args) == 1 {
		cmd.Help()
		return nil
	}

	configFile := configAccess.GetDefaultFilename()
	if configAccess.IsExplicitFile() {
		configFile = configAccess.GetExplicitFile()
	}

	_, ok := config.Contexts[name]
	if !ok {
		return fmt.Errorf("cannot delete context %s, not in %s", name, configFile)
	}

	if config.CurrentContext == name {
		fmt.Fprint(errOut, "warning: this removed your active context, use \"kubectl config use-context\" to select a different one\n")
	}

	delete(config.Contexts, name)

	if err := clientcmd.ModifyConfig(configAccess, *config, true); err != nil {
		return err
	}

	fmt.Fprintf(out, "deleted context %s from %s\n", name, configFile)

	return nil
}
