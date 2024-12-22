package functionality

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func DeleteDeploy(namespace string, deployment string) {

    // Get the user's home directory
    homeDir, err := os.UserHomeDir()
    if err != nil {
        panic(fmt.Sprintf("Failed to get home directory: %v", err))
    }

    // Load kubeconfig
    kubeconfig := filepath.Join(homeDir, ".kube", "config")
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        panic(fmt.Sprintf("Failed to load kubeconfig: %v", err))
    }

    // Create Kubernetes clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(fmt.Sprintf("Failed to create Kubernetes client: %v", err))
    }

    // Define the deployment
    deletePolicy := metav1.DeletePropagationForeground
	err = clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), deployment, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to delete deployment: %v\n", err)
		os.Exit(1)
	}

    fmt.Println("Deployment deleted successfully!")
}

