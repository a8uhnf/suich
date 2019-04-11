package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func ChangeKubectl() *cobra.Command {
	var version string
	cmd := &cobra.Command{
		Use:   "kubectl",
		Short: "Update to provided kubectl version.",
		Long:  "Update to provided kubectl version. Kubectl version must be provided. now by default machine type set to amd64",
		Run: func(cmd *cobra.Command, args []string) {
			k := &Kubectl{}
			err := k.Validate(cmd)
			if err != nil {
				return
			}
			err = k.Downloads(cmd)
			if err != nil {
				return
				panic(err)
			}
		},
	}
	cmd.Flags().StringVarP(&version, "version", "v", "v1.9.0", "kubectl valid version")
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

func (k *Kubectl) Downloads(cmd *cobra.Command) error {
	v := cmd.Flag("version").Value.String()
	// o, err := os.
	var out bytes.Buffer
	c := exec.Command("uname")
	c.Stdout = &out
	err := c.Run()
	if err != nil {
		return err
	}
	o := string(out.Bytes())
	url := fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/%s/bin/%s/amd64/kubectl", v, strings.ToLower(strings.Trim(o, "\n")))
	fmt.Println("--------", url)
	c = exec.Command("curl", "-LO", url)
	// c.Stdout = &out
	if _, err := os.Stat("./kubectl"); err != nil {
		err = c.Run()
		if err != nil {
			return err
		}
	}

	c = exec.Command("chmod", "a+wx", "./kubectl")
	err = c.Run()
	if err != nil {
		return err
	}
	out.Reset()
	c = exec.Command("which", "kubectl")
	c.Stdout = &out
	err = c.Run()
	fmt.Println("**********************", string(out.Bytes()))
	log.Printf("Moving to: %s", string(out.Bytes()))
	if err != nil {
		return err
	}
	dest := string(out.Bytes())
	if dest != "/usr/local/bin/kubectl" {
		dest = "/usr/local/bin/kubectl"
	} else {
		out.Reset()
		c = exec.Command("chmod", "a+wx", dest)
		c.Stderr = &out
		err = c.Run()
		if err != nil {
			fmt.Println(string(out.Bytes()))
			panic(err)
		}
	}

	fmt.Println(dest)
	// c = exec.Command("mv", "./kubectl", strings.Trim(dest, "\n"))
	// c.Stderr = &out
	// err = c.Run()
	err = os.Rename("kubectl", dest)
	if err != nil {
		panic(err)
	}
	fmt.Println("**********************", string(out.Bytes()))
	if err != nil {
		return err
	}
	return nil
}
