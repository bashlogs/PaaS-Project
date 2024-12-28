package functionality

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateNamespace(clientset *kubernetes.Clientset, name string) {

	if GetNamespace(clientset, name) != nil {
		fmt.Println("Namespace already existed")
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
}

func GetNamespace(clientset *kubernetes.Clientset, namespace string) *v1.Namespace {
	ns, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
        fmt.Printf("Failed to get namespace: %v", err)
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

func RoleBinding(clientset *kubernetes.Clientset, namespace string){
	role := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "lower-role",
			Namespace: namespace,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""}, // Core API group
				Resources: []string{"pods", "services"},
				Verbs:     []string{"create", "delete", "get", "list", "watch"},
			},
			{
				APIGroups: []string{"apps"}, // Apps API group
				Resources: []string{"deployments"},
				Verbs:     []string{"create", "delete", "get", "list", "watch"},
			},
		},
	}

	_, err := clientset.RbacV1().Roles(namespace).Create(context.TODO(), role, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create Role: %v", err)
	}
	fmt.Println("Role created successfully")

	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nothing",
			Namespace: namespace,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "username",  // Change to "ServiceAccount" if needed
				Name:      "noting", // Replace with your user or service account name
				Namespace: namespace,      // Required for ServiceAccount
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "example-role",
		},
	}
	
	// Create the RoleBinding
	_, err = clientset.RbacV1().RoleBindings(namespace).Create(context.TODO(), roleBinding, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create RoleBinding: %v", err)
	}
	fmt.Println("RoleBinding created successfully")
}