package kubernetes

import (
	"context"
	"errors"

	"k8s.io/client-go/kubernetes"

	"github.com/bashlogs/PaaS_Project/api/api"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// func CreateDeployment(w http.ResponseWriter, r *http.Request) {
// 	var param = api.KubernetesInfo{}
// 	err := json.NewDecoder(r.Body).Decode(&param)
// 	if err != nil {
// 		log.Error("Error decoding request body:", err)
// 		api.InternalErrorHandler(w)
// 		return
// 	}

// 	err = CheckDeployment(param.FrontendURL, param.BackendURL)
// 	if err != nil {
// 		log.Error("Error checking deployment:", err)
// 		api.InternalErrorHandler(w)
// 		return
// 	}

// 	fmt.Println("Check Deployment check passed")

//     database, err := tools.ConnectToDatabase()
//     if err != nil {
//         log.Error("Database connection error: ", err)
//         api.InternalErrorHandler(w)
//         return
//     }

// 	rows, err := database.DB.Query("update namespace set frontend_image = $2, backend_image = $3 where namespace_id = $1;", param.Id, param.FrontendURL, param.BackendURL)

// 	if err != nil {
// 		api.RequestErrorHandler(w, errors.New("no workspaces found"))
// 		return
// 	}

// 	defer rows.Close()

// 	fmt.Println("Update query executed successfully")

// 	// Check for the user namespace if not create a new one

// 	database, err = tools.ConnectToDatabase()
//     if err != nil {
//         log.Error("Database connection error: ", err)
//         api.InternalErrorHandler(w)
//         return
//     }

// 	if rows.Next() {
// 		err := rows.Scan(&param.Username, &param.Namespace)
// 		if err != nil {
// 			api.RequestErrorHandler(w, errors.New("failed to scan result"))
// 			return
// 		}
// 	} else {
// 		api.RequestErrorHandler(w, errors.New("no workspaces found"))
// 		return
// 	}

// 	ok, err := Check_namespace(param.Username, param.Namespace)

// 	if err != nil {
// 		log.Error("Error checking namespace:", err)
// 		api.InternalErrorHandler(w)
// 		return
// 	}

// 	if ok {
// 		log.Error("Namespace already exists")
// 	}
// 	// assign resources to the namespace

// 	// do the deployment

// 	// create ingress service

// 	// Finally, send a response back to the client

//     var workspaces []api.Workspace
//     for rows.Next() {
//         var workspace api.Workspace
//         err := rows.Scan(&workspace.Id, &workspace.Name, &workspace.IsActive)
//         if err != nil {
//             log.Error("Error scanning row:", err)
//             api.InternalErrorHandler(w)
//             return
//         }
//         workspace.Endpoint = fmt.Sprintf("http://localhost:8080/%s", workspace.Name)
//         workspaces = append(workspaces, workspace)
//     }

//     if err = rows.Err(); err != nil {
//         log.Error("Error iterating rows:", err)
//         api.InternalErrorHandler(w)
//         return
//     }

//     fmt.Println("Workspace data: ", workspaces)

// 	w.Header().Set("Content-Type", "application/json")
// 	err = json.NewEncoder(w).Encode(workspaces)
// 	if err != nil {
// 		log.Error("Error encoding dashboard response:", err)
//         http.Error(w, "Internal server error", http.StatusInternalServerError)
//         return
//     }
// }



func GetKubeNamespaceInfo(clientset *kubernetes.Clientset, namespace string) (*v1.Namespace, error) {
	ns, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		return nil, err
    }
	return ns, nil
}

func CreateNamespace(clientset *kubernetes.Clientset, name string) error {
	ns, err := GetKubeNamespaceInfo(clientset, name)
	if err == nil && ns != nil {
		return errors.New("Namespace already existed")
	}

    namespace := &v1.Namespace{
        ObjectMeta: metav1.ObjectMeta{
            Name: name,
        },
    }

	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
    if err != nil {
		err := errors.New("Failed to create namespace")
		return err
    }

	return nil
}

// func GetResourceQuota(namespace string) (string, error) {
// 	fmt.Println("Resource quota found: ", namespace)
// 	err := errors.New("resource quota not found")
// 	return "", err;
// }

func SetResourceQuota(clientset *kubernetes.Clientset, namespace string) error {
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

func RollbackDeployment(clientset *kubernetes.Clientset, namespace string) error {
	err := clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}

func SetConfigMap(clientset *kubernetes.Clientset, name string, configMaps []api.ConfigMaps) (*v1.ConfigMap, error) {
	configMap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Data: make(map[string]string),
	}

	for _, config := range configMaps {
		configMap.Data[config.Key] = config.Value
	}

	_, err := clientset.CoreV1().ConfigMaps("default").Create(context.TODO(), configMap, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return configMap, nil

}

// func SetDeployment(clientset *kubernetes.Clientset, Namespace string, name string, image string, port int32, configMap *v1.ConfigMap) error {
// 	// Create a deployment
// 	deployment := &v1.Deployment{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: name,
// 		},
// 		Spec: v1.DeploymentSpec{
// 			Replicas: int32Ptr(1),
// 			Selector: &metav1.LabelSelector{
// 				MatchLabels: map[string]string{"app": name},
// 			},
// 			Template: v1.PodTemplateSpec{
// 				ObjectMeta: metav1.ObjectMeta{
// 					Labels: map[string]string{"app": name},
// 				},
// 				Spec: v1.PodSpec{
// 					Containers: []v1.Container{
// 						{
// 							Name:  name,
// 							Image: image,
// 							Ports: []v1.ContainerPort{{ContainerPort: port}},
// 							EnvFrom: []v1.EnvFromSource{
// 								{
// 									ConfigMapRef: &v1.ConfigMapEnvSource{
// 										LocalObjectReference: v1.LocalObjectReference{Name: configMap.Name},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	_, err := clientset.AppsV1().Deployments(Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
//     if err != nil {
//         fmt.Fprintf(os.Stderr, "Failed to create backend deployment: %v\n", err)
// 		return err;
//     }

// 	return nil;
// }


// func SetService(clientset *kubernetes.Clientset, name string, namespace string, port int32) (int32, error) {
	
// 	servicesClient := clientset.CoreV1().Services(namespace)

// 	// Check if the service exists
// 	_, err := servicesClient.Get(context.TODO(), name, metav1.GetOptions{})
// 	if err == nil {
// 		// Service exists, delete it
// 		fmt.Printf("Service %s already exists in namespace %s. Deleting it...\n", name, namespace)
// 		err = servicesClient.Delete(context.TODO(), name, metav1.DeleteOptions{})
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Failed to delete existing service: %v\n", err)
// 			return err
// 	} else if !errors.IsNotFound(err) {
// 	} else if !apierrors.IsNotFound(err) {
// 		// Other error (not "not found")
// 		fmt.Fprintf(os.Stderr, "Error checking for existing service: %v\n", err)
// 		return err
// 	}
	
	
// 	service := &v1.Service{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: name,
// 		},
// 		Spec: v1.ServiceSpec{
// 			Type: v1.ServiceTypeClusterIP,
// 			Ports: []v1.ServicePort{
// 				{
// 					Port: port,
// 				},
// 			},
// 			Selector: map[string]string{"app": name},
// 		},
// 	}

// 	svc, err := clientset.CoreV1().Services("default").Create(context.TODO(), service, metav1.CreateOptions{})
// 	if err != nil {
// 		return 0, err
// 	}

// 	return svc.Spec.Port, nil

// }