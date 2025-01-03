package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bashlogs/kubernetes-go/functionality"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	
	homeDir, err := os.UserHomeDir()
    if err != nil {
        panic(fmt.Sprintf("Failed to get home directory: %v", err))
    }

    kubeconfig := filepath.Join(homeDir, ".kube", "config")
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        panic(fmt.Sprintf("Failed to load kubeconfig: %v", err))
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(fmt.Sprintf("Failed to create Kubernetes client: %v", err))
    }
	
	// delete the deployment
	// functionality.DeleteDeploy("default", "nginx-deployment")
	// functionality.createDeploy()

	// functionality.CreateNamespace(clientset, "mayur123")
	// functionality.DeleteNamespace(clientset, "mayur123")

	// functionality.ResourceAllocation(clientset, "mayur")

	// functionality.CreateNamespace(clientset, "project")

	functionality.VolumeClaim(clientset, "project")
}

