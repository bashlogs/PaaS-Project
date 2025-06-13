package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bashlogs/PaaS_Project/api/api"
	"github.com/bashlogs/PaaS_Project/api/internal/kubernetes"
	log "github.com/sirupsen/logrus"
	k8sclient "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateDeployment2(w http.ResponseWriter, r *http.Request) {
	var deployments map[string]api.DeploymentInfo

	err := json.NewDecoder(r.Body).Decode(&deployments)
	if err != nil {
		log.Error("Error decoding request body:", err)
		api.InternalErrorHandler(w)
		return
	}

	if len(deployments) == 0 {
		log.Error("No deployment data provided")
		api.ClientErrorHandler(w)
		return
	}

	homeDir, err := os.UserHomeDir()
    if err != nil {
        panic(fmt.Sprintf("Failed to get home directory: %v", err))
    }

    kubeconfig := filepath.Join(homeDir, ".kube", "config")
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        panic(fmt.Sprintf("Failed to load kubeconfig: %v", err))
    }

    clientset, err := k8sclient.NewForConfig(config)
    if err != nil {
        panic(fmt.Sprintf("Failed to create Kubernetes client: %v", err))
    }
	
	if dep, ok := deployments["backend_deployment"]; ok {
		name := dep.Username + "-" + dep.Namespace
		err := kubernetes.CreateNamespace(clientset, name)

		if err != nil {
			log.Error("Error creating namespace:", err)
			api.KubernetesErrorHandler(w, err)
			return
		}

		err = kubernetes.SetResourceQuota(clientset, name)
		if err != nil {
			log.Error("Error creating resource quota:", err)
			kubernetes.RollbackDeployment(clientset, name)
			api.KubernetesErrorHandler(w, err)
			return
		}

		backendConfigMap, err := kubernetes.SetConfigMap(clientset, name, "backend-config", dep.ConfigMaps)
		if err != nil {
			log.Error("Error creating backend configmap:", err)
			kubernetes.RollbackDeployment(clientset, name)
			api.KubernetesErrorHandler(w, err)
			return
		}

		err = kubernetes.SetDeployment(clientset, name, dep.Name, dep.Image, dep.Port, backendConfigMap)
		if err != nil {
			log.Error("Error creating backend deployment:", err)
			kubernetes.RollbackDeployment(clientset, name)
			api.KubernetesErrorHandler(w, err)
			return
		}

		port, err := kubernetes.SetService(clientset, dep.Name, name, dep.Port)
		if err != nil {
			log.Error("Error creating backend service:", err)
			kubernetes.RollbackDeployment(clientset, name)
			api.KubernetesErrorHandler(w, err)
			return
		}

		if dep, ok := deployments["frontend_deployment"]; ok {
			dep.ConfigMaps = append(dep.ConfigMaps, api.ConfigMaps{
				Key:   "backend-service",
				Value: fmt.Sprintf("%d", port),
			})
		}

		fmt.Println("Deployment created successfully")
		fmt.Println(dep)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Deployment created successfully"))
		return
	}

	// if dep, ok := deployments["frontend_deployment"]; ok {
	// 	name := dep.Username + "-" + dep.Namespace
	// 	err := creaeteNamespace(name)
	// 	if err != nil {
	// 		log.Error("Error creating namespace:", err)
	// 		api.KubernetesErrorHandler(w, err)
	// 		return
	// 	}

	// 	err := SetResourceQuota(name)
	// 	if err != nil {
	// 		log.Error("Error creating resource quota:", err)
	// 		RollbackDeployment(name)
	// 		api.KubernetesErrorHandler(w, err)
	// 		return
	// 	}

	// 	configmap, err := SetConfigMap("frontend-config", dep.ConfigMaps)
	// 	if err != nil {
	// 		log.Error("Error creating frontend configmap:", err)
	// 		RollbackDeployment(name)
	// 		api.KubernetesErrorHandler(w, err)
	// 		return
	// 	}

	// 	err = SetDeployment(dep.Name, dep.Image, dep.Port, configmap)
	// 	if err != nil {
	// 		log.Error("Error creating frontend deployment:", err)
	// 		RollbackDeployment(name)
	// 		api.KubernetesErrorHandler(w, err)
	// 		return
	// 	}

	// 	port, err := SetService(dep.Name, dep.Port)
	// 	if err != nil {
	// 		log.Error("Error creating frontend service:", err)
	// 		RollbackDeployment(name)
	// 		api.KubernetesErrorHandler(w, err)
	// 		return
	// 	}

	// 	err = SetIngress(dep.Name, port)
	// 	if err != nil {
	// 		log.Error("Error creating ingress service:", err)
	// 		RollbackDeployment(name)
	// 		api.KubernetesErrorHandler(w, err)
	// 		return
	// 	}
	// }
}
