package main

import (
    "context"
    "fmt"
    "os"

    appsv1 "k8s.io/api/apps/v1"
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/util/retry"
)

func main() {
    // Create a Kubernetes client
    kubeconfig := filepath.Join(homeDir, ".kube", "config")
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        panic(err.Error())
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    // Create the deployment object
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

    // Create the deployment
    retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
        _, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
        return err
    })
    if retryErr != nil {
        fmt.Fprintf(os.Stderr, "Failed to create deployment: %v\n", retryErr)
        os.Exit(1)
    }

    fmt.Println("Deployment created successfully!")
}

func int32Ptr(i int32) *int32 { return &i }

