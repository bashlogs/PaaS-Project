"use client"

import { useState, useEffect } from "react"
import { useParams } from "next/navigation"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { DockerDeploy } from "@/components/workspace/docker-deploy"
import { DeploymentInfo } from "@/components/workspace/deployment-info"
import { getUserWorkspaces, type Workspace } from "@/lib/api/workspaces"

export default function WorkspacePage() {
  const { id } = useParams()
  const [workspace, setWorkspace] = useState<Workspace | null>(null)
  const [deploymentFlow, setDeploymentFlow] = useState<string[]>([])
  const [logs, setLogs] = useState<string[]>([])

  useEffect(() => {
    getUserWorkspaces().then((workspaces) => {
      const currentWorkspace = workspaces.find((w) => w.id === id)
      if (currentWorkspace) {
        setWorkspace(currentWorkspace)
      }
    })
  }, [id])

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
    <Card>
      <CardHeader>
        <CardTitle>{workspace.name}</CardTitle>
        <CardDescription>Deploy your application using Docker</CardDescription>
      </CardHeader>
      <CardContent className="space-y-6">
        <DockerDeploy onDeploy={handleDeploy} />
        {(deploymentFlow.length > 0 || logs.length > 0) && (
          <DeploymentInfo deploymentFlow={deploymentFlow} logs={logs} />
        )}
      </CardContent>
    </Card>
  )
}

