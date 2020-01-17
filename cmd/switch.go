package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"

	go_prompt "github.com/c-bata/go-prompt"

	"github.com/a8uhnf/suich/pkg/utils"
)

var kubeConfigPath = filepath.Join(os.Getenv("HOME"), ".kube", "config")
var prompt bool

func SwitchCmd() *cobra.Command {
	cc := &cobra.Command{
		Use:   "switch",
		Short: "To switch context use this command",
		PreRun: func(cmd *cobra.Command, args []string) {
			if prompt {
				fmt.Println("-------------")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var selectedCtx string
			log.Println("Starting reading config file....")
			ctxs, err := readKubeConfigFile()
			if err != nil {
				log.Fatalln(err)
			}
			if !prompt {
				selectedCtx, err = utils.RunPrompt(ctxs, "Select Context")
				if err != nil {
					log.Fatalln(err)
				}
			} else {
				sugges := []go_prompt.Suggest{}
				for _, k := range ctxs {
					sugges = append(sugges, go_prompt.Suggest{k, ""})
				}
				go_prompt.New(executor, completer)
			}

			c := exec.Command("kubectl", "config", "use-context", selectedCtx)
			var out bytes.Buffer
			c.Stdout = &out
			err = c.Run()
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
	cc.Flags().BoolVarP(&prompt, "prompt", "p", false, "run switch in prompt mode")
	return cc
}
func executor(in string) {
	fmt.Println(in)
}

func completer(go_prompt.Document) []go_prompt.Suggest {

	return []go_prompt.Suggest{}
}

// readKubeConfigFile reads kubeconfig and return context list
func readKubeConfigFile() ([]string, error) {
	file, err := ioutil.ReadFile(kubeConfigPath)
	if err != nil {
		return nil, err
	}
	cfg := KubectlConfig{}
	log.Println("Successfully read kube-config...")

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
