"use client";

import { createContext, useContext, useState, useCallback, ReactNode } from "react";
import { 
  getUserWorkspaces, 
  createWorkspace,
  updateWorkspaceStatus,
  updateWorkspaceEndpoint,
  deleteWorkspace,
  type Workspace 
} from "@/lib/api/workspaces";

// Define the context value type
type WorkspaceContextType = {
  workspaces: Workspace[];
  fetchWorkspaces: (forceRefresh?: boolean) => Promise<void>;
  createNewWorkspace: (name: string, endpoint: string, username: string) => Promise<void>;
  toggleWorkspaceStatus: (id: string, currentStatus: boolean) => Promise<void>;
  updateEndpoint: (id: string, newEndpoint: string) => Promise<void>;
  removeWorkspace: (id: string) => Promise<void>;
};

// Create the context with a default value
const WorkspaceContext = createContext<WorkspaceContextType | undefined>(undefined);

// Provider props type
type WorkspaceProviderProps = {
  children: ReactNode;
};

export const WorkspaceProvider = ({ children }: WorkspaceProviderProps) => {
  const [workspaces, setWorkspaces] = useState<Workspace[]>([]);

  const fetchWorkspaces = useCallback(async (forceRefresh = false) => {
    try {
      const data = await getUserWorkspaces(forceRefresh);
      setWorkspaces(data);
    } catch (error) {
      console.error("Error fetching workspaces:", error);
    }
  }, []);

  const createNewWorkspace = useCallback(async (name: string, endpoint: string, username: string) => {
    try {
      await createWorkspace(name, endpoint, username);
      await fetchWorkspaces(true);
    } catch (error) {
      console.error("Failed to create workspace:", error);
      throw error;
    }
  }, [fetchWorkspaces]);

  const toggleWorkspaceStatus = useCallback(async (id: string, currentStatus: boolean) => {
    try {
      await updateWorkspaceStatus(id, !currentStatus);
      await fetchWorkspaces(true);
    } catch (error) {
      console.error("Failed to update workspace status:", error);
      throw error;
    }
  }, [fetchWorkspaces]);

  const updateEndpoint = useCallback(async (id: string, newEndpoint: string) => {
    try {
      await updateWorkspaceEndpoint(id, newEndpoint);
      await fetchWorkspaces(true);
    } catch (error) {
      console.error("Failed to update workspace endpoint:", error);
      throw error;
    }
  }, [fetchWorkspaces]);

  const removeWorkspace = useCallback(async (id: string) => {
    try {
      await deleteWorkspace(id);
      await fetchWorkspaces(true);
    } catch (error) {
      console.error("Failed to delete workspace:", error);
      throw error;
    }
  }, [fetchWorkspaces]);

  return (
    <WorkspaceContext.Provider 
      value={{ 
        workspaces, 
        fetchWorkspaces, 
        createNewWorkspace, 
        toggleWorkspaceStatus, 
        updateEndpoint, 
        removeWorkspace 
      }}
    >
      {children}
    </WorkspaceContext.Provider>
  );
};

// Custom hook to use the workspace context
export const useWorkspace = (): WorkspaceContextType => {
  const context = useContext(WorkspaceContext);
  if (context === undefined) {
    throw new Error("useWorkspace must be used within a WorkspaceProvider");
  }
  return context;
};