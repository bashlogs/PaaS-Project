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

func ServiceAccount(clientset *kubernetes.Clientset, namespace string){
	service_account := "sva-1"
	createServiceAccount(clientset, namespace, service_account)
	createClusterRole(clientset)
	createGlobalClusterRoleBinding(clientset, namespace, service_account)
}

func createServiceAccount(clientset *kubernetes.Clientset, namespace string, service_account string) {
	serviceAccount := &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: service_account,
		},
	}
	_, err := clientset.CoreV1().ServiceAccounts(namespace).Create(context.TODO(), serviceAccount, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create ServiceAccount: %v", err)
	}
	fmt.Println("ServiceAccount created successfully in namespace:", namespace)
}

func createClusterRole(clientset *kubernetes.Clientset) {
	clusterRole := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: "global-access-role",
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""}, // Core API group
				Resources: []string{"pods", "services"},
				Verbs:     []string{"create", "delete", "get", "list", "watch"},
			},
			{
				APIGroups: []string{"apps"},
				Resources: []string{"deployments"},
				Verbs:     []string{"create", "delete", "get", "list", "watch"},
			},
			{
				APIGroups: []string{""}, // Core API group
				Resources: []string{"configmaps", "secrets"},
				Verbs:     []string{"get", "list"},
			},
		},
	}

	_, err := clientset.RbacV1().ClusterRoles().Create(context.TODO(), clusterRole, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create ClusterRole: %v", err)
	}
	fmt.Println("ClusterRole created successfully")
}


func createGlobalClusterRoleBinding(clientset *kubernetes.Clientset, namespace string, serviceAccountName string) {
	clusterRoleBinding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "global-access-binding",
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      serviceAccountName,
				Namespace: namespace,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "global-access-role",
		},
	}

	_, err := clientset.RbacV1().ClusterRoleBindings().Create(context.TODO(), clusterRoleBinding, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create ClusterRoleBinding: %v", err)
	}
	fmt.Println("ClusterRoleBinding created successfully")
}