"use client"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"

interface DeploymentInfoProps {
  deploymentFlow: string[]
  logs: string[]
}

export function DeploymentInfo({ deploymentFlow, logs }: DeploymentInfoProps) {
  return (
    <Tabs defaultValue="flow" className="w-full">
      <TabsList>
        <TabsTrigger value="flow">Deployment Flow</TabsTrigger>
        <TabsTrigger value="logs">Logs</TabsTrigger>
      </TabsList>
      <TabsContent value="flow">
        <div className="bg-gray-100 p-4 rounded-md">
          {deploymentFlow.map((step, index) => (
            <div key={index} className="mb-2">
              <span className="font-bold">{index + 1}.</span> {step}
            </div>
          ))}
        </div>
      </TabsContent>
      <TabsContent value="logs">
        <pre className="bg-black text-green-400 p-4 rounded-md overflow-x-auto">{logs.join("\n")}</pre>
      </TabsContent>
    </Tabs>
  )
}

