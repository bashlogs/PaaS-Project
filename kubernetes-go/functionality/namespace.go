package functionality

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateNamespace(clientset *kubernetes.Clientset, Username string, Namespace string) {
	name := Username + "-" + Namespace

	if GetNamespace(clientset, name) != nil {
		fmt.Println("Namespace already existed")
		// ServiceAccount(clientset, name)
		return
	}

    namespace := &v1.Namespace{
        ObjectMeta: metav1.ObjectMeta{
            Name: name,
        },
    }

	_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
    if err != nil {
        panic(fmt.Sprintf("Failed to create namespace: %v", err))
    }

    fmt.Println("Namespace created successfully")

	// ServiceAccount(clientset, name)
}

func GetNamespace(clientset *kubernetes.Clientset, namespace string) *v1.Namespace {
	ns, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		return nil
    }
	return ns
}

func DeleteNamespace(clientset *kubernetes.Clientset, namespace string){
	default_namespaces := []string{"default", "kube-node-lease", "kube-public", "kube-system"}

	for _, names := range default_namespaces {
		if names == namespace {
			panic(fmt.Sprint("Failed to delete default namespace"))
		}
	}

	err := clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{})
	if err != nil {
		panic(fmt.Sprintf("Failed to default namespace: %v", err))
	}

	fmt.Println("Namespace deleted successfully")
}