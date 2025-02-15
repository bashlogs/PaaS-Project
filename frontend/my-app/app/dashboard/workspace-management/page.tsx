"use client"

import { useState, useEffect } from "react"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import {
  getUserWorkspaces,
  createWorkspace,
  renameWorkspace,
  deleteWorkspace,
  type Workspace,
} from "@/lib/api/workspaces"

export default function WorkspaceManagementPage() {
  const [workspaces, setWorkspaces] = useState<Workspace[]>([])
  const [newWorkspaceName, setNewWorkspaceName] = useState("")
  const [editingWorkspace, setEditingWorkspace] = useState<Workspace | null>(null)

  useEffect(() => {
    loadWorkspaces()
  }, [])

  const loadWorkspaces = async () => {
    const loadedWorkspaces = await getUserWorkspaces()
    setWorkspaces(loadedWorkspaces)
  }

  const handleCreateWorkspace = async () => {
    if (newWorkspaceName.trim()) {
      await createWorkspace(newWorkspaceName.trim())
      setNewWorkspaceName("")
      loadWorkspaces()
    }
  }

  const handleRenameWorkspace = async (id: string, newName: string) => {
    await renameWorkspace(id, newName)
    setEditingWorkspace(null)
    loadWorkspaces()
  }

  const handleDeleteWorkspace = async (id: string) => {
    if (confirm("Are you sure you want to delete this workspace?")) {
      await deleteWorkspace(id)
      loadWorkspaces()
    }
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle>Workspace Management</CardTitle>
        <CardDescription>Create, rename, and delete your workspaces</CardDescription>
      </CardHeader>
      <CardContent>
        <div className="space-y-6">
          <div className="space-y-2">
            <Label htmlFor="new-workspace">Create New Workspace</Label>
            <div className="flex space-x-2">
              <Input
                id="new-workspace"
                value={newWorkspaceName}
                onChange={(e) => setNewWorkspaceName(e.target.value)}
                placeholder="Enter workspace name"
              />
              <Button onClick={handleCreateWorkspace}>Create</Button>
            </div>
          </div>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Workspace Name</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {workspaces.map((workspace) => (
                <TableRow key={workspace.id}>
                  <TableCell>
                    {editingWorkspace?.id === workspace.id ? (
                      <Input
                        value={editingWorkspace.name}
                        onChange={(e) => setEditingWorkspace({ ...editingWorkspace, name: e.target.value })}
                      />
                    ) : (
                      workspace.name
                    )}
                  </TableCell>
                  <TableCell>
                    {editingWorkspace?.id === workspace.id ? (
                      <div className="space-x-2">
                        <Button onClick={() => handleRenameWorkspace(workspace.id, editingWorkspace.name)}>Save</Button>
                        <Button variant="outline" onClick={() => setEditingWorkspace(null)}>
                          Cancel
                        </Button>
                      </div>
                    ) : (
                      <div className="space-x-2">
                        <Button variant="outline" onClick={() => setEditingWorkspace(workspace)}>
                          Rename
                        </Button>
                        <Button variant="destructive" onClick={() => handleDeleteWorkspace(workspace.id)}>
                          Delete
                        </Button>
                      </div>
                    )}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>
      </CardContent>
    </Card>
  )
}

