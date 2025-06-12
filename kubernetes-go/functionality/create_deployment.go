package functionality

import (
	"context"
	"errors"
	"fmt"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

func CreateNamespace2(clientset *kubernetes.Clientset, name string) error{

	if GetNamespace(clientset, name) != nil {
		err := errors.New("Namespace already existed")
		return err
	}

    namespace := &v1.Namespace{
        ObjectMeta: metav1.ObjectMeta{
            Name: name,
        },
    }

	_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
    if err != nil {
		err := errors.New("Failed to create namespace")
		return err
    }

    fmt.Println("Namespace created successfully")

	return nil
}

func ResourceAllocation2(clientset *kubernetes.Clientset, namespace string) error {
	_, err := clientset.CoreV1().ResourceQuotas(namespace).Get(context.TODO(), "low-quota", metav1.GetOptions{})
	if err == nil {
		return nil;
	}
	Quota := &v1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "low-quota",
			Namespace: namespace,
		},
		Spec: v1.ResourceQuotaSpec{
			Hard: v1.ResourceList{
				v1.ResourceCPU:             resource.MustParse("1"),
				v1.ResourceMemory:          resource.MustParse("1Gi"),
				v1.ResourcePods:            resource.MustParse("5"),
				v1.ResourceRequestsStorage: resource.MustParse("2Gi"),
			},
		},
	}

	_, err = clientset.CoreV1().ResourceQuotas(namespace).Create(context.TODO(), Quota, metav1.CreateOptions{})
	if err != nil {
		err := errors.New("Failed to create resource quota")
		return err
	}

	return nil
}

func backenddeploy(clientset *kubernetes.Clientset, BackendUrl string, Port int, Namespace string) error {
	deployment := &appsv1.Deployment{
        ObjectMeta: metav1.ObjectMeta{
            Name: "backend-deployment",
        },
        Spec: appsv1.DeploymentSpec{
            Replicas: int32Ptr(1),
            Selector: &metav1.LabelSelector{
                MatchLabels: map[string]string{
                    "app": "backend-deployment",
                },
            },
            Template: corev1.PodTemplateSpec{
                ObjectMeta: metav1.ObjectMeta{
                    Labels: map[string]string{
                        "app": "backend-deployment",
                    },
                },
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{
                        {
                            Name:  "backend-deployment",
                            Image: BackendUrl,
							EnvFrom: []corev1.EnvFromSource{
								{
									Prefix: "BACKEND_URL",
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "backend",
										},
									},
								},
							},
                            Ports: []corev1.ContainerPort{
                                {
                                    ContainerPort: int32(Port),
                                },
                            },
                        },
                    },
                },
            },
        },
    }

    // Create the deployment in the "default" namespace
    _, err := clientset.AppsV1().Deployments(Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create backend deployment: %v\n", err)
		return err;
    }

	return nil;
}

func frontenddeploy(clientset *kubernetes.Clientset, FrontendUrl string, Port int, Namespace string) error {
	deployment := &appsv1.Deployment{
        ObjectMeta: metav1.ObjectMeta{
            Name: "frontend-deployment",
        },
        Spec: appsv1.DeploymentSpec{
            Replicas: int32Ptr(1),
            Selector: &metav1.LabelSelector{
                MatchLabels: map[string]string{
                    "app": "frontend-deployment",
                },
            },
            Template: corev1.PodTemplateSpec{
                ObjectMeta: metav1.ObjectMeta{
                    Labels: map[string]string{
                        "app": "frontend-deployment",
                    },
                },
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{
                        {
                            Name:  "frontend-deployment",
                            Image: FrontendUrl,
							EnvFrom: []corev1.EnvFromSource{
								{
									Prefix: "FRONTEND_URL",
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "frontend",
										},
									},
								},
							},
                            Ports: []corev1.ContainerPort{
                                {
                                    ContainerPort: int32(Port),
                                },
                            },
                        },
                    },
                },
            },
        },
    }

    // Create the deployment in the "default" namespace
    _, err := clientset.AppsV1().Deployments(Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create frontend deployment: %v\n", err)
    }

	return nil;
}


func CreateDeploy2(clientset *kubernetes.Clientset, FrontendUrl string, BackendUrl string, Namespace string) error {
	
	fmt.Println("Started Backend Deployment")
	err := backenddeploy(clientset, BackendUrl, 5000, Namespace)
	if err != nil {
		fmt.Println("Error creating backend deployment:", err)
	}

	fmt.Println("Backend Deployment created successfully")

	fmt.Println("Started Frontend Deployment")

	err = frontenddeploy(clientset, FrontendUrl, 3000, Namespace)
	if err != nil {
		fmt.Println("Error creating frontend deployment:", err)
		return err
	}

	return nil;
}

func CreateConfigMap(clientset *kubernetes.Clientset, Namespace string, Name string, configs map[string]string) error {
	
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      Name,
			Namespace: Namespace,
		},
		Data: configs,
	}

	_, err := clientset.CoreV1().ConfigMaps(Namespace).Create(context.TODO(), configMap, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create configmap: %v\n", err)
		return err
	}

	return nil
}

func CreateNodePortService(clientset *kubernetes.Clientset, namespace string, name string, port int) error {
	servicesClient := clientset.CoreV1().Services(namespace)

	// Check if the service exists
	_, err := servicesClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err == nil {
		// Service exists, so delete it first
		fmt.Printf("Service %s already exists in namespace %s. Deleting it...\n", name, namespace)
		err = servicesClient.Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete existing service: %v\n", err)
			return err
		} else if !apierrors.IsNotFound(err) {
			// Some other error occurred
			fmt.Fprintf(os.Stderr, "Error checking for existing service: %v\n", err)
			return err
		}
	}

	// Define the new service
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Ports: []corev1.ServicePort{
				{
					Port:     int32(port),
					NodePort: int32(30000),
					Protocol: corev1.ProtocolTCP,
				},
			},
			Selector: map[string]string{
				"app": name,
			},
		},
	}

	// Create the new service
	_, err = servicesClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create service: %v\n", err)
		return err
	}

	fmt.Println("Service created successfully.")
	return nil
}


func CreateClusterIPService(clientset *kubernetes.Clientset, namespace string, name string, port int) error {
	servicesClient := clientset.CoreV1().Services(namespace)

	// Check if the service exists
	_, err := servicesClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err == nil {
		// Service exists, delete it
		fmt.Printf("Service %s already exists in namespace %s. Deleting it...\n", name, namespace)
		err = servicesClient.Delete(context.TODO(), name, metav1.DeleteOptions{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete existing service: %v\n", err)
			return err
		}
	} else if !apierrors.IsNotFound(err) {
		// Other error (not "not found")
		fmt.Fprintf(os.Stderr, "Error checking for existing service: %v\n", err)
		return err
	}

	// Define the ClusterIP service
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeClusterIP,
			Ports: []corev1.ServicePort{
				{
					Port:       int32(port),
					TargetPort: intstr.FromInt(30000),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: map[string]string{
				"app": name,
			},
		},
	}

	// Create the new service
	_, err = servicesClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create service: %v\n", err)
		return err
	}

	fmt.Println("ClusterIP service created successfully.")
	return nil
}


func CreateIngressService(clientset *kubernetes.Clientset, namespace, name string, port int) error {
	ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Annotations: map[string]string{
				"traefik.ingress.kubernetes.io/router.entrypoints": "web",
			},
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: "bashlogs.local",
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     "/" + namespace,
									PathType: func() *networkingv1.PathType {
										pathType := networkingv1.PathTypePrefix
										return &pathType
									}(),
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: "frontend-service",
											Port: networkingv1.ServiceBackendPort{
												Number: int32(port),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := clientset.NetworkingV1().Ingresses(namespace).Create(context.TODO(), ingress, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Failed to create ingress: %v\n", err)
		return err
	}
	fmt.Println("Ingress created successfully")
	return nil
}


