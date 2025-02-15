export interface Workspace {
  id: string
  name: string
}

let mockWorkspaces: Workspace[] = [
  { id: "workspace-1", name: "Personal Project" },
  { id: "workspace-2", name: "Team Alpha" },
  { id: "workspace-3", name: "Client X" },
]

export async function getUserWorkspaces(): Promise<Workspace[]> {
  await new Promise((resolve) => setTimeout(resolve, 500))
  return mockWorkspaces
}

export async function createWorkspace(name: string): Promise<Workspace> {
  await new Promise((resolve) => setTimeout(resolve, 500))
  const newWorkspace: Workspace = {
    id: `workspace-${mockWorkspaces.length + 1}`,
    name,
  }
  mockWorkspaces.push(newWorkspace)
  return newWorkspace
}

export async function renameWorkspace(id: string, newName: string): Promise<Workspace> {
  await new Promise((resolve) => setTimeout(resolve, 500))
  const workspace = mockWorkspaces.find((w) => w.id === id)
  if (!workspace) {
    throw new Error("Workspace not found")
  }
  workspace.name = newName
  return workspace
}

export async function deleteWorkspace(id: string): Promise<void> {
  await new Promise((resolve) => setTimeout(resolve, 500))
  mockWorkspaces = mockWorkspaces.filter((w) => w.id !== id)
}

