export interface Workspace {
  id: string;
  name: string;
  isActive: boolean;
  endpoint: string;
}

let cachedWorkspaces: Workspace[] | null = null;

async function fetchWorkspaces(forceRefresh: boolean = false): Promise<Workspace[]> {
  if (!forceRefresh && cachedWorkspaces !== null) {
    console.log("Returning cached workspaces:", cachedWorkspaces);
    return cachedWorkspaces; // Return cached data if available and no refresh is required
  }

  try {
    console.log("Fetching workspaces from API...");
    const response = await fetch("http://localhost:8000/api/workspaces", {
      method: "GET",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });

    if (!response.ok) {
      let errorData;
      try {
        errorData = await response.json();
      } catch {
        errorData = { message: "Unknown error" };
      }
      console.error("API Error:", response.status, response.statusText, errorData);
      throw new Error(`Error ${response.status}: ${errorData?.message || "Failed to fetch workspaces"}`);
    }

    const data = await response.json();

    // Transform the API response to match the Workspace interface
    cachedWorkspaces = data.map((item: any) => ({
      id: String(item.namespace_id),
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
  return fetchWorkspaces(forceRefresh);
}

// Function to invalidate the cache when changes are made
function invalidateWorkspaceCache() {
  cachedWorkspaces = null;
}

// Modify these functions to invalidate cache when updating or deleting a workspace
export async function createWorkspace(name: string, endpoint: string, username: string): Promise<Workspace> {
  const response = await fetch("http://127.0.0.1:8000/api/workspaces", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    credentials: "include",
    body: JSON.stringify({ name, endpoint, username }),
  });

  if (!response.ok) {
    throw new Error("Failed to create workspace");
  }

  const newWorkspace: Workspace = await response.json();
  invalidateWorkspaceCache(); // Invalidate cache after creation
  return newWorkspace;
}

export async function updateWorkspaceStatus(id: string, isActive: boolean): Promise<Workspace> {
  const response = await fetch(`http://127.0.0.1:8000/api/workspaces/${id}`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ isActive }),
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Failed to update workspace status");
  }

  const updatedWorkspace: Workspace = await response.json();
  invalidateWorkspaceCache(); // Invalidate cache after update
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

  const updatedWorkspace: Workspace = await response.json();
  invalidateWorkspaceCache(); // Invalidate cache after update
  return updatedWorkspace;
}

export async function deleteWorkspace(id: string): Promise<void> {
  const response = await fetch(`http://localhost:8000/api/workspaces/${id}`, {
    method: "DELETE",
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Failed to delete workspace");
  }

  invalidateWorkspaceCache(); // Invalidate cache after deletion
}
