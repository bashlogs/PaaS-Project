export interface Workspace {
  id: string;
  name: string;
  isActive: boolean;
  endpoint: string;
}

let cachedWorkspaces: Workspace[] | null = null;

async function fetchWorkspaces(forceRefresh: boolean): Promise<Workspace[]> {
  if (!forceRefresh && cachedWorkspaces) {
    console.log("Returning cached workspaces:", cachedWorkspaces);
    return cachedWorkspaces;
  }

  try {
    console.log("Fetching workspaces from API...");
    const response = await fetch("http://localhost:8000/api/workspaces", {
      method: "GET",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });

    if (!response.ok) {
      const errorData = await response.json();
      console.error("API Error:", response.status, response.statusText, errorData);
      throw new Error(`Error ${response.status}: ${errorData?.message || "Failed to fetch workspaces"}`);
    }

    const data = await response.json();
    cachedWorkspaces = data.map((item: any) => ({
      id: item.namespace_id,
      name: item.namespace,
      isActive: item.active,
      endpoint: item.endpoint,
    }));

    console.log("Updated workspace cache:", cachedWorkspaces);
    return cachedWorkspaces;
  } catch (error) {
    console.error("Fetch failed:", error);
    return [];
  }
}

export async function getUserWorkspaces(forceRefresh: boolean = false): Promise<Workspace[]> {
  return await fetchWorkspaces(forceRefresh);
}

// Function to invalidate the cache when changes are made
function invalidateWorkspaceCache() {
  cachedWorkspaces = null;
}

// Modify these functions to invalidate cache when updating or deleting a workspace
export async function createWorkspace(name: string, endpoint: string, username: string): Promise<Workspace> {
  const response = await fetch("http://localhost:8000/api/workspaces", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    credentials: "include",
    body: JSON.stringify({ name, endpoint, username }),
  });

  if (!response.ok) {
    throw new Error("Failed to create workspace");
  }

  const newWorkspace = await response.json();
  fetchWorkspaces(true); 
  const workspaceInfo: Workspace = {
    id: newWorkspace.namespace_id,
    name: newWorkspace.namespace,
    isActive: newWorkspace.active,
    endpoint: newWorkspace.endpoint,
  };

  return workspaceInfo; 
}

export async function updateWorkspaceStatus(id: string, isActive: boolean): Promise<Workspace> {
  const response = await fetch(`http://localhost:8000/api/workspaces_status`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    credentials: "include",
    body: JSON.stringify({ id, isActive }),
  });

  if (!response.ok) {
    throw new Error("Failed to update workspace status");
  }

  const updatedWorkspace = await response.json();
  return updatedWorkspace;
}

export async function updateWorkspaceEndpoint(id: string, endpoint: string): Promise<Workspace> {
  const response = await fetch(`http://localhost:8000/api/workspaces/${id}`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ endpoint }),
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Failed to update workspace endpoint");
  }

  const updatedWorkspace = await response.json();
  invalidateWorkspaceCache(); // Invalidate cache after update
  return updatedWorkspace;
}

export async function deleteWorkspace(id: string): Promise<void> {
  const response = await fetch(`http://localhost:8000/api/workspaces?id=${id}`, {
    method: "DELETE",
    credentials: "include", 
  });

  if (!response.ok) {
    throw new Error("Failed to delete workspace");
  }
  const output = fetchWorkspaces(true); 
  console.log(output);
  const deleteworkspace = await response.json();
  return deleteworkspace
}
