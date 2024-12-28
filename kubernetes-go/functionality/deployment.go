package functionality

import (
	"context"
	"fmt"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateDeploy(clientset *kubernetes.Clientset) {
    // Define the deployment
    deployment := &appsv1.Deployment{
        ObjectMeta: metav1.ObjectMeta{
            Name: "nginx-deployment",
        },
        Spec: appsv1.DeploymentSpec{
            Replicas: int32Ptr(1),
            Selector: &metav1.LabelSelector{
                MatchLabels: map[string]string{
                    "app": "nginx",
                },
            },
            Template: corev1.PodTemplateSpec{
                ObjectMeta: metav1.ObjectMeta{
                    Labels: map[string]string{
                        "app": "nginx",
                    },
                },
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{
                        {
                            Name:  "nginx",
                            Image: "nginx:latest",
                            Ports: []corev1.ContainerPort{
                                {
                                    ContainerPort: 80,
                                },
                            },
                        },
                    },
                },
            },
        },
    }

    // Create the deployment in the "default" namespace
    _, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create deployment: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Deployment created successfully!")
}

func DeleteDeploy(clientset *kubernetes.Clientset, namespace string, deployment string) {

    // Define the deployment
    deletePolicy := metav1.DeletePropagationForeground
	err := clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), deployment, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to delete deployment: %v\n", err)
		os.Exit(1)
	}

    fmt.Println("Deployment deleted successfully!")
}


// Helper function to get a pointer to an int32
func int32Ptr(i int32) *int32 { return &i }

