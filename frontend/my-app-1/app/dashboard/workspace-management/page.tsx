"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "@/store/store";
import { fetchUserData } from "@/store/userSlice";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Badge } from "@/components/ui/badge"
import { useWorkspace } from "@/components/contexts/WorkspaceContext";

const WorkspaceManagement = () => {
  const router = useRouter();
  const dispatch = useDispatch<AppDispatch>();
  const { userData, isLoading } = useSelector((state: RootState) => state.user);
  
  const { 
    workspaces, 
    fetchWorkspaces, 
    createNewWorkspace, 
    toggleWorkspaceStatus, 
    updateEndpoint, 
    removeWorkspace 
  } = useWorkspace();
  
  const [newWorkspaceName, setNewWorkspaceName] = useState("")
  const [newWorkspaceEndpoint, setNewWorkspaceEndpoint] = useState("")

  useEffect(() => {
    dispatch(fetchUserData())
      .unwrap()
      .then(userData => {
        console.log("User data fetched successfully for workspace management:", userData);
      })
      .catch(error => {
        console.error("Failed to fetch user data:", error);
        router.push("/login");
      });
  }, [dispatch, router]);
  
  useEffect(() => {
    fetchWorkspaces();
  }, [fetchWorkspaces]);

  useEffect(() => {
    if (userData) {
      const username = userData.username;
      setNewWorkspaceEndpoint(`http://localhost:8000/${username}/${newWorkspaceName}`);
    }
  }, [newWorkspaceName, userData]);

  const handleCreateWorkspace = async () => {
    if (userData && newWorkspaceName.trim() && newWorkspaceEndpoint.trim()) {
      try {
        await createNewWorkspace(
          newWorkspaceName.trim(), 
          newWorkspaceEndpoint.trim(), 
          userData.username
        );
        setNewWorkspaceName("");
        setNewWorkspaceEndpoint("");
      } catch (error) {
        console.error("Failed to create workspace:", error);
      }
    }
  };

  const handleToggleStatus = async (id: string, currentStatus: boolean) => {
    await toggleWorkspaceStatus(id, currentStatus);
  }

  const handleUpdateEndpoint = async (id: string, newEndpoint: string) => {
    await updateEndpoint(id, newEndpoint);
  }

  const handleDeleteWorkspace = async (id: string) => {
    if (confirm(`Are you sure you want to delete this workspace?`)) {
      try {
        await removeWorkspace(id);
      } catch (error) {
        console.error("Failed to delete workspace:", error);
      }
    }
  };

  if (isLoading) return <p>Loading...</p>;
  
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
  );
};

export default WorkspaceManagement;