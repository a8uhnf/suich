package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var kubeConfigPath = filepath.Join(os.Getenv("HOME"), ".kube", "config")

func SwitchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "switch",
		Short: "To switch context use this command",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Starting reading config file....")
			ctxs, err := readKubeConfigFile()
			if err != nil {
				log.Fatalln(err)
			}

			selectedCtx, err := runPrompt(ctxs)
			if err != nil {
				log.Fatalln(err)
			}

			
		},
	}
}

func readKubeConfigFile() ([]string, error) {
	file, err := ioutil.ReadFile(kubeConfigPath)
	if err != nil {
		return nil, err
	}
	cfg := KubectlConfig{}
	log.Println("Succeessfully read kube-config...")

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}
	ret := []string{}

	for _, v := range cfg.Contexts {
		ret = append(ret, v.Name)
	}
	return ret, nil
}

type KubectlConfig struct {
	Kind           string                    `json:"kind"`
	ApiVersion     string                    `json:"apiVersion"`
	CurrentContext string                    `json:"current-context"`
	Clusters       []*KubectlClusterWithName `json:"clusters"`
	Contexts       []*KubectlContextWithName `json:"contexts"`
	Users          []*KubectlUserWithName    `json:"users"`
}

type KubectlClusterWithName struct {
	Name    string         `json:"name"`
	Cluster KubectlCluster `json:"cluster"`
}

type KubectlCluster struct {
	Server                   string `json:"server,omitempty"`
	CertificateAuthorityData []byte `json:"certificate-authority-data,omitempty"`
}

type KubectlContextWithName struct {
	Name    string         `json:"name"`
	Context KubectlContext `json:"context"`
}

type KubectlContext struct {
	Cluster string `json:"cluster"`
	User    string `json:"user"`
}

type KubectlUserWithName struct {
	Name string      `json:"name"`
	User KubectlUser `json:"user"`
}

type KubectlUser struct {
	ClientCertificateData []byte `json:"client-certificate-data,omitempty"`
	ClientKeyData         []byte `json:"client-key-data,omitempty"`
	Password              string `json:"password,omitempty"`
	Username              string `json:"username,omitempty"`
	Token                 string `json:"token,omitempty"`
}

func runPrompt(ctxs []string) (string, error) {
	list := []string{}
	for _, v := range ctxs {
		list = append(list, v)
	}
	prompt := promptui.Select{
		Label: "Select context",
		Items: list,
	}
	_, res, err := prompt.Run()
	if err != nil {
		return "", nil
	}
	return res, nil
}