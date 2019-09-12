package cmd

import (
	"bytes"
	"errors"
	"os"
	"strings"

	"github.com/a8uhnf/suich/pkg/utils"
	"github.com/spf13/cobra"
)

const (
	podInfoNameTitle = "NAME"
)

var (
	follow = false
)

// GetLogsCmd builds the logs cobra command for suich
func GetLogsCmd() *cobra.Command {
	logsCMD := &cobra.Command{
		Use:   "logs",
		Short: "Get logs for a certain pod",
		Long:  "Prompts the user with a list of pods names to display the logs",
		RunE:  getLogs,
	}
	logsCMD.Flags().StringP("namespace", "n", "", "The namespace to work on")
	logsCMD.Flags().BoolVarP(&follow, "follow", "f", false, "Watch the logs")

	return logsCMD
}

func getLogs(cmd *cobra.Command, args []string) error {

	client := getKubernetesClient()

	var namespace string

	nList, err := getNamespaceNames(client)
	if err != nil {
		namespace, err = cmd.Flags().GetString("namespace")
		if err != nil {
			return err
		}
	} else {
		namespace, err = utils.RunPrompt(nList, "Select Namespace")
		if err != nil {
			return err
		}
	}

	if namespace == "" {
		return errors.New("Must provide namespace flag as you do not have access to list namespaces")
	}

	var podsBfr bytes.Buffer

	if err := utils.ExecCommand(&podsBfr, "kubectl", "get", "pods", "-n", namespace); err != nil {
		return err
	}

	pns := readAllPods(podsBfr)

	pod, err := utils.RunPrompt(pns, "Select Pod")

	if err != nil {
		return err
	}

	follow, err := cmd.Flags().GetBool("follow")
	if err != nil {
		return err
	}
	if follow {
		if err := utils.ExecCommand(os.Stdout, "kubectl", "logs", pod, "-n", namespace, "-f"); err != nil {
			return err
		}
	} else {
		if err := utils.ExecCommand(os.Stdout, "kubectl", "logs", pod, "-n", namespace); err != nil {
			return err
		}
	}

	return nil
}

func readAllPods(out bytes.Buffer) []string {
	podInfos := strings.Split(out.String(), "\n")

	podNames := []string{}

	for _, pi := range podInfos {
		info := strings.Split(pi, " ")

		if info[0] == podInfoNameTitle || info[0] == "" {
			continue
		}

		podNames = append(podNames, info[0])
	}

	return podNames

}
