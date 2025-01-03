package functionality

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)



func VolumeClaim(clientset *kubernetes.Clientset, namespace string){
	pvc := "pvc-1"
	createPVC(clientset, namespace, pvc)
}


func createPVC(clientset *kubernetes.Clientset, namespace string, pvcName string) {
    pvc := &v1.PersistentVolumeClaim{
        ObjectMeta: metav1.ObjectMeta{
            Name: pvcName,
        },
        Spec: v1.PersistentVolumeClaimSpec{
            AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce}, // Access mode for the volume (single node access)
            Resources: v1.VolumeResourceRequirements{
                Requests: v1.ResourceList{
                    v1.ResourceStorage: resource.MustParse("1Gi"),
                },
            },
			StorageClassName: ptrToString("standard"),
        },
    }

    // Attempt to create the PVC
    _, err := clientset.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
    if err != nil {
        log.Fatalf("Failed to create PersistentVolumeClaim: %v", err) // Handle error appropriately
    }
    
    fmt.Println("PersistentVolumeClaim created successfully in namespace:", namespace)
}

func ptrToString(s string) *string {
    return &s
}