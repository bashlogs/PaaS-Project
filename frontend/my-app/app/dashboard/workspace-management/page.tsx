"use client"

import { useState, useEffect, useRef } from "react"
import { useRouter } from "next/navigation"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Badge } from "@/components/ui/badge"
import { useDispatch, useSelector } from "react-redux";
import { RootState, AppDispatch } from "@/store/store";
import { fetchUserData } from "@/store/userSlice";
import {
  getUserWorkspaces,
  createWorkspace,
  updateWorkspaceStatus,
  updateWorkspaceEndpoint,
  deleteWorkspace,
  type Workspace,
} from "@/lib/api/workspaces"

export default function WorkspaceManagementPage() {
  const dispatch = useDispatch<AppDispatch>();
  const { userData, isLoading, error } = useSelector((state: RootState) => state.user);
  const router = useRouter();

  const [workspaces, setWorkspaces] = useState<Workspace[]>([])
  const [newWorkspaceName, setNewWorkspaceName] = useState("")
  const [newWorkspaceEndpoint, setNewWorkspaceEndpoint] = useState("")
    
  useEffect(() => {
    console.log("Dispatching fetchUserData...");
    dispatch(fetchUserData()).catch((error) => console.error("Failed to fetch user data:", error));
  }, [dispatch]);

  const loadWorkspaces = async () => {
    try {
      const workspaces = await getUserWorkspaces(true);
      setWorkspaces(workspaces);
    } catch (error) {
      console.error("Failed to fetch workspaces:", error);
    }
  }
  
  useEffect(() => {
    loadWorkspaces();
  }, []);

  useEffect(() => {
    if (error === "Unauthorized") {
      console.log("Unauthorized error detected, clearing auth state...");
      
      document.cookie.split(";").forEach(cookie => {
        const [name] = cookie.split("=");
        document.cookie = `${name.trim()}=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/;`;
      });
      
      if (typeof window !== 'undefined') {
        localStorage.removeItem('authToken');
      }
  
      setTimeout(() => {
        router.push("/login");
      }, 100);
    }
  }, [error, router]);

  useEffect(() => {
    if (userData) {
      const username = userData.username;
      setNewWorkspaceEndpoint(`http://localhost:8000/${username}/${newWorkspaceName}`);
    }
  }, [newWorkspaceName, userData]);
  
  const handleCreateWorkspace = async () => {
    if (userData && newWorkspaceName.trim() && newWorkspaceEndpoint.trim()) {
      try {
        const newWorkspace = await createWorkspace(newWorkspaceName.trim(), newWorkspaceEndpoint.trim(), userData.username);
        await loadWorkspaces();
        setNewWorkspaceName("");
        setNewWorkspaceEndpoint("");
      } catch (error) {
        console.error("Failed to create workspace:", error);
      }
    }
  };

  const handleToggleStatus = async (id: string, currentStatus: boolean) => {
    await updateWorkspaceStatus(id, !currentStatus)
    loadWorkspaces()
  }

  const handleUpdateEndpoint = async (id: string, newEndpoint: string) => {
    await updateWorkspaceEndpoint(id, newEndpoint)
    loadWorkspaces()
  }

  const handleDeleteWorkspace = async (id: string) => {
    if (confirm(`Are you sure you want to delete this ${id} workspace?`)) {
      try {
        await deleteWorkspace(id);
        await loadWorkspaces();
        // Update the state by filtering out the deleted workspace
        // setWorkspaces(prevWorkspaces => prevWorkspaces.filter(workspace => workspace.id !== id));
      } catch (error) {
        console.error("Failed to delete workspace:", error);
      }
    }
  };
  

  // if (isLoading) return <div>Loading...</div>;
  // if (error === "Unauthorized") return <p>Error: {error}</p>;

  return (
    <Card>
      <CardHeader>
        <CardTitle>Workspace Management</CardTitle>
        <CardDescription>Create and manage your workspaces</CardDescription>
      </CardHeader>
      <CardContent className="space-y-6">
        <div className="space-y-2">
          <h3 className="text-lg font-medium">Create New Workspace</h3>
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="new-workspace-name">Workspace Name</Label>
              <Input
                id="new-workspace-name"
                value={newWorkspaceName}
                onChange={(e) => setNewWorkspaceName(e.target.value)}
                placeholder="Enter workspace name"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="new-workspace-endpoint">Endpoint URL</Label>
              <Input
                id="new-workspace-endpoint"
                value={newWorkspaceEndpoint}
                onChange={(e) => setNewWorkspaceEndpoint(e.target.value)}
                placeholder="https://example.com"
              />
            </div>
          </div>
          <Button onClick={handleCreateWorkspace} className="mt-2">
            Create Workspace
          </Button>
        </div>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Workspace Name</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>Endpoint</TableHead>
              <TableHead>Actions</TableHead>
              <TableHead>Delete</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {workspaces && workspaces.length > 0 ? (
              workspaces.map((workspace) => (
                <TableRow key={workspace.id}>
                  <TableCell>{workspace.name}</TableCell>
                  <TableCell>
                    <Badge variant={workspace.isActive ? "success" : "destructive"}>
                      {workspace.isActive ? "Active" : "Inactive"}
                    </Badge>
                  </TableCell>
                  <TableCell>
                    <Input
                      value={workspace.endpoint}
                      onChange={(e) => handleUpdateEndpoint(workspace.id, e.target.value)}
                    />
                  </TableCell>
                  <TableCell>
                    <Button
                      variant={workspace.isActive ? "destructive" : "default"}
                      onClick={() => handleToggleStatus(workspace.id, workspace.isActive)}
                      className="w-full"
                    >
                      {workspace.isActive ? "Deactivate" : "Activate"}
                    </Button>
                  </TableCell>
                  <TableCell>
                    <Button variant="destructive" onClick={() => handleDeleteWorkspace(workspace.id)} className="w-full">
                      Delete
                    </Button>
                  </TableCell>
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell colSpan={5}>No workspaces available</TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  )
}