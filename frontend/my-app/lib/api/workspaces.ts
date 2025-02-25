export interface Workspace {
  id: string
  name: string
  isActive: boolean
  endpoint: string
}

let mockWorkspaces: Workspace[] = [
  { id: "workspace-1", name: "Personal Project", isActive: true, endpoint: "https://personal.example.com" },
  { id: "workspace-2", name: "Team Alpha", isActive: true, endpoint: "https://alpha.example.com" },
  { id: "workspace-3", name: "Client X", isActive: false, endpoint: "https://clientx.example.com" },
]

export async function getUserWorkspaces(): Promise<Workspace[]> {
  await new Promise((resolve) => setTimeout(resolve, 500))
  return mockWorkspaces
}

export async function createWorkspace(name: string, endpoint: string): Promise<Workspace> {
  await new Promise((resolve) => setTimeout(resolve, 500))
  const newWorkspace: Workspace = {
    id: `workspace-${mockWorkspaces.length + 1}`,
    name,
    isActive: true,
    endpoint,
  }
  mockWorkspaces.push(newWorkspace)
  return newWorkspace
}

export async function updateWorkspaceStatus(id: string, isActive: boolean): Promise<Workspace> {
  await new Promise((resolve) => setTimeout(resolve, 500))
  const workspace = mockWorkspaces.find((w) => w.id === id)
  if (!workspace) {
    throw new Error("Workspace not found")
  }
  workspace.isActive = isActive
  return workspace
}

export async function updateWorkspaceEndpoint(id: string, endpoint: string): Promise<Workspace> {
  await new Promise((resolve) => setTimeout(resolve, 500))
  const workspace = mockWorkspaces.find((w) => w.id === id)
  if (!workspace) {
    throw new Error("Workspace not found")
  }
  workspace.endpoint = endpoint
  return workspace
}

export async function deleteWorkspace(id: string): Promise<void> {
  await new Promise((resolve) => setTimeout(resolve, 500))
  mockWorkspaces = mockWorkspaces.filter((w) => w.id !== id)
}

