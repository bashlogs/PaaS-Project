package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bashlogs/kubernetes-go/functionality"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	
	homeDir, err := os.UserHomeDir()
    if err != nil {
        panic(fmt.Sprintf("Failed to get home directory: %v", err))
    }

    kubeconfig := filepath.Join(homeDir, ".kube", "config")
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        panic(fmt.Sprintf("Failed to load kubeconfig: %v", err))
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(fmt.Sprintf("Failed to create Kubernetes client: %v", err))
    }
	
	// delete the deployment
	// functionality.DeleteDeploy("default", "nginx-deployment")
	// functionality.createDeploy()
	
	// functionality.CreateNamespace(clientset, Username, Namespace)
	
	// functionality.DeleteNamespace(clientset, "mayur123")

	// functionality.ResourceAllocation(clientset, "mayur")

	// functionality.CreateNamespace(clientset, "project")

	// functionality.VolumeClaim(clientset, "project")

	// To Create Namespace

	Username := "mayur"
	Namespace := "123"
	name := Username + "-" + Namespace

	err = functionality.CreateNamespace2(clientset, name)

	if err != nil {
		fmt.Println("Error creating namespace:", err)
	} else {
		fmt.Println("Namespace created successfully:", name)
	}


	// assign resources to the namespace

	err = functionality.ResourceAllocation2(clientset, name)

	if err != nil {
		fmt.Println("Error creating resource allocation:", err)
		return
	} else {
		fmt.Println("Resource allocation created successfully:", name)
	}

	// create configmaps
	frontend_configs := map[string]string{
		"config.js":  `window._env_ = {
			REACT_APP_BACKEND_URL: "http://backend.local"
		};`,
	}

	backend_configs := map[string]string{
		"MONGO_URI": "",
		"PORT": "5000",
	}

	err = functionality.CreateConfigMap(clientset, name, "frontend", frontend_configs)

	if err != nil {
		fmt.Println("Error creating frontend configmap:", err)
	} else {
		fmt.Println("Frontend Configmap created successfully")
	}

	err = functionality.CreateConfigMap(clientset, name, "backend", backend_configs)

	if err != nil {
		fmt.Println("Error creating backend configmap:", err)
	} else {
		fmt.Println("Backend Configmap created successfully")
	}
	
	// do the deployment

	err = functionality.CreateDeploy2(clientset, "nginx:latest", "nginx:latest", name)

	if err != nil {
		fmt.Println("Error creating deployment:", err)
		return
	} else {
		fmt.Println("Deployment created successfully")
	}

	// Creating services for deployments

	err = functionality.CreateNodePortService(clientset, name, "frontend-service", 80)

	if err != nil {
		fmt.Println("Error creating frontend service:", err)
		return
	} else {
		fmt.Println("Frontend Service created successfully")	
	}

	err = functionality.CreateClusterIPService(clientset, name, "backend-service", 80)

	if err != nil {
		fmt.Println("Error creating frontend service:", err)
		return
	} else {
		fmt.Println("Backend Service created successfully")	
	}

	// create ingress service

	err = functionality.CreateIngressService(clientset, name, "frontend-service", 30000)
	if err != nil {
		fmt.Println("Error creating ingress service:", err)
		return
	} else {
		fmt.Println("Ingress Service created successfully")
	}

	// Finally, send a response back to the client


}

