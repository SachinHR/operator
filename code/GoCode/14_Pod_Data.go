//export KUBERNETES_SERVICE_HOST=localhost
//export KUBERNETES_SERVICE_PORT=51113
//export KUBERNETES_SERVICE_PORT=50247
//export KUBECONFIG=~/.kube/config

package main

import (
    "context"
    "fmt"
    "os"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
)

func main() {
    kubeconfigPath := os.Getenv("KUBECONFIG")
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
    if err != nil {
        panic(err)
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err)
    }

    pods, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "POD-NAME", metav1.GetOptions{})
    if err != nil {
        panic(err)
    }

	fmt.Println(pods) 
}
