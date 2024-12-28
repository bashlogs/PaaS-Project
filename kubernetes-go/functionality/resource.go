package functionality

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ResourceAllocation(clientset *kubernetes.Clientset, namespace string) {

	CreateNamespace(clientset, namespace)

	Quota := &v1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
            Name: "low-quota",
			Namespace: namespace,
        },
		Spec: v1.ResourceQuotaSpec{
			Hard: v1.ResourceList{
				v1.ResourceCPU: resource.MustParse("1"),         // Limit to 2 CPUs
				v1.ResourceMemory: resource.MustParse("1Gi"),       // Limit to 4Gi of memory
				v1.ResourcePods: resource.MustParse("5"),        // Limit to 10 pods
				v1.ResourceRequestsStorage: resource.MustParse("2Gi"),    // Limit to 20Gi storage
			},
			ScopeSelector: &v1.ScopeSelector{
				MatchExpressions: []v1.ScopedResourceSelectorRequirement{
					{
						Operator: v1.ScopeSelectorOpIn,
						ScopeName: v1.ResourceQuotaScopePriorityClass,
						Values: []string{"low"},
					},
				},
			},
		},
	}

	_, err := clientset.CoreV1().ResourceQuotas(namespace).Create(context.TODO(), Quota, metav1.CreateOptions{})
    if err != nil {
        panic(fmt.Sprintf("Failed to create namespace: %v", err))
    }

    fmt.Println("Resouce quota create successfully")
}