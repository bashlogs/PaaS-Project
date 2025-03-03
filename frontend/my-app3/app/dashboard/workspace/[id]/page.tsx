"use client"

import { useState, useEffect } from "react"
import { useParams } from "next/navigation"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { DockerDeploy } from "@/components/workspace/docker-deploy"
import { DeploymentInfo } from "@/components/workspace/deployment-info"
import { InfoSteps } from "@/components/workspace/info-steps"
import { getUserWorkspaces, type Workspace } from "@/lib/api/workspaces"

export default function WorkspacePage() {
  const { id } = useParams() as { id: string };
  const [workspace, setWorkspace] = useState<Workspace | null>(null)
  const [step, setStep] = useState<"info" | "deploy">("info")
  const [deploymentFlow, setDeploymentFlow] = useState<string[]>([])
  const [logs, setLogs] = useState<string[]>([])

  useEffect(() => {
    if (!id) return;
  
    const controller = new AbortController();
  
    getUserWorkspaces()
      .then((workspaces) => {
        if (!workspaces || !Array.isArray(workspaces)) {
          console.error("Invalid workspace response:", workspaces);
          return;
        }
        const currentWorkspace = workspaces.find((w) => w.id === id);
        if (currentWorkspace) {
          setWorkspace(currentWorkspace);
        } else {
          console.error("Workspace not found for ID:", id);
        }
      })
      .catch((error) => {
        if (error.name !== "AbortError") {
          console.error("Error fetching workspaces:", error);
        }
      });
  
    return () => controller.abort();
  }, [id]);
  

  const handleNext = () => {
    setStep("deploy")
  }

  const handleDeploy = (frontendUrl: string, backendUrl: string) => {
    // Simulated deployment process
    setDeploymentFlow([
      "Fetching Dockerfiles",
      "Building frontend container",
      "Building backend container",
      "Pushing containers to registry",
      "Updating deployment configuration",
      "Deploying to Kubernetes cluster",
      "Waiting for services to be ready",
      "Deployment complete",
    ])

    setLogs([
      "INFO: Deployment started",
      `INFO: Frontend Dockerfile URL: ${frontendUrl}`,
      `INFO: Backend Dockerfile URL: ${backendUrl}`,
      "INFO: Building frontend container...",
      "INFO: Frontend container built successfully",
      "INFO: Building backend container...",
      "INFO: Backend container built successfully",
      "INFO: Pushing containers to registry...",
      "INFO: Containers pushed successfully",
      "INFO: Updating deployment configuration...",
      "INFO: Deploying to Kubernetes cluster...",
      "INFO: Waiting for services to be ready...",
      "INFO: Deployment complete",
      "INFO: Application is now accessible",
    ])
  }

  if (!workspace) {
    return <div>Loading...</div>
  }

  return (
    <div className="space-y-6">
      <h2 className="text-3xl font-bold tracking-tight">{workspace.name}</h2>

      {step === "info" ? (
        <InfoSteps onNext={handleNext} />
      ) : (
        <Card>
          <CardHeader>
            <CardTitle>Deploy Your Application</CardTitle>
            <CardDescription>Provide your Docker image URLs to start the deployment</CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            <DockerDeploy onDeploy={handleDeploy} />
            {(deploymentFlow.length > 0 || logs.length > 0) && (
              <DeploymentInfo deploymentFlow={deploymentFlow} logs={logs} />
            )}
          </CardContent>
        </Card>
      )}
    </div>
  )
}

