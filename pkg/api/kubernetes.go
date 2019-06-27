package api

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/a8uhnf/suich/pkg/utils"
)
// GetKubernetesClient generate kubernetes local client
func GetKubernetesClient() kubernetes.Interface {
	var kubeClient kubernetes.Interface
	kubeClient = utils.GetClientOutOfCluster()
	return kubeClient
}
// GetNamespaceNames generate list of ns name from k8s nsList
func GetNamespaceNames(namespaces *v1.NamespaceList)[]string {
	ret := []string{}
	for _, v := range namespaces.Items {
		ret = append(ret, v.ObjectMeta.Name)
	}
	fmt.Println(ret)
	return ret
}