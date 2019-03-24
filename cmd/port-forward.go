package cmd

import (

	// k8s "k8s.io/kubernetes"

	"fmt"

	"github.com/spf13/cobra"

	"github.com/a8uhnf/suich/pkg/api"
	"github.com/a8uhnf/suich/pkg/utils"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func PortForward() *cobra.Command {
	return &cobra.Command{
		Use:   "pf",
		Short: "port-forward kubernetes pod.",
		Run: func(cmd *cobra.Command, args []string) {
			kClient := getKubernetesClient()
			// if err != nil {
			// 	log.Fatalln(err)
			// }

			nList, err := getNamespaceNames(kClient)
			if err != nil {
				panic(err)
			}
			n, err := runPrompt(nList)
			if err != nil {
				panic(err)
			}
			fmt.Println(n)
			pList, err := getSpecificNSPods(kClient, n)
			if err != nil {
				panic(err)
			}
			p, err := runPrompt(pList)
			if err != nil {
				panic(err)
			}
			fmt.Println(p)
		},
	}
}

func getKubernetesClient() kubernetes.Interface {
	var kubeClient kubernetes.Interface
	// _, err := rest.InClusterConfig()
	kubeClient = utils.GetClientOutOfCluster()
	// if err != nil {

	// } else {
	// 	kubeClient = utils.GetClient()
	// }
	return kubeClient
}

func getSpecificNSPods(c kubernetes.Interface, ns string) ([]string, error) {
	pods, err := c.CoreV1().Pods(ns).List(meta_v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	p := api.PodList{}
	for _, v := range pods.Items {
		fmt.Println(v.ObjectMeta.Name)
		p.Pods = append(p.Pods, v.ObjectMeta.Name)
	}

	return p.Pods, nil
}
func getNamespaceNames(c kubernetes.Interface) ([]string, error) {
	ns, err := c.CoreV1().Namespaces().List(meta_v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	nsList := api.NamespaceList{}
	for _, v := range ns.Items {
		nsList.Namespaces = append(nsList.Namespaces, v.ObjectMeta.Name)
	}
	return nsList.Namespaces, nil
}
