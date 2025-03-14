import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"

interface InfoStepsProps {
  onNext: () => void
}

export function InfoSteps({ onNext }: InfoStepsProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Getting Started with Deployment</CardTitle>
        <CardDescription>Follow these steps to prepare your project for deployment</CardDescription>
      </CardHeader>
      <CardContent className="space-y-6">
        <ol className="list-decimal list-inside space-y-4">
          <li>Ensure your project has a valid Dockerfile for both frontend and backend.</li>
          <li>Make sure all environment variables are properly set in your project.</li>
          <li>Commit and push your latest changes to your version control system.</li>
          <li>Have your Docker image repository URLs ready (for both frontend and backend).</li>
          <li>Prepare any additional configuration files required for your deployment.</li>
        </ol>
        <Button onClick={onNext} className="w-full">
          Next: Configure Deployment
        </Button>
      </CardContent>
    </Card>
  )
}

