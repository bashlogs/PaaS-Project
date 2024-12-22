package main

import "github.com/bashlogs/kubernetes-go/functionality"

func main() {
	
	// delete the deployment
	functionality.DeleteDeploy("default", "nginx-deployment")
}

