"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"

interface DockerDeployProps {
  onDeploy: (frontendUrl: string, backendUrl: string) => void
}

export function DockerDeploy({ onDeploy }: DockerDeployProps) {
  const [frontendUrl, setFrontendUrl] = useState("")
  const [backendUrl, setBackendUrl] = useState("")

  const handleDeploy = () => {
    onDeploy(frontendUrl, backendUrl)
  }

  return (
    <div className="space-y-4">
      <div className="space-y-2">
        <Label htmlFor="frontend-url">Frontend Dockerfile URL</Label>
        <Input
          id="frontend-url"
          placeholder="https://example.com/frontend/Dockerfile"
          value={frontendUrl}
          onChange={(e) => setFrontendUrl(e.target.value)}
        />
      </div>
      <div className="space-y-2">
        <Label htmlFor="backend-url">Backend Dockerfile URL</Label>
        <Input
          id="backend-url"
          placeholder="https://example.com/backend/Dockerfile"
          value={backendUrl}
          onChange={(e) => setBackendUrl(e.target.value)}
        />
      </div>
      <Button onClick={handleDeploy} className="w-full">
        Deploy
      </Button>
    </div>
  )
}

