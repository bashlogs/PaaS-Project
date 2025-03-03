// "use client"

// import { useState, useEffect, useRef } from "react"
// import { useRouter } from "next/navigation"
// import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
// import { Button } from "@/components/ui/button"
// import { Input } from "@/components/ui/input"
// import { Label } from "@/components/ui/label"
// import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
// import { Badge } from "@/components/ui/badge"
// import {
//   getUserWorkspaces,
//   createWorkspace,
//   updateWorkspaceStatus,
//   updateWorkspaceEndpoint,
//   deleteWorkspace,
//   type Workspace,
// } from "@/lib/api/workspaces"

// export default function WorkspaceManagementPage() {

//   const [workspaces, setWorkspaces] = useState<Workspace[]>([])
//   const [newWorkspaceName, setNewWorkspaceName] = useState("")
//   const [newWorkspaceEndpoint, setNewWorkspaceEndpoint] = useState("")
  
//   const loadWorkspaces = async () => {
//     try {
//       const workspaces = await getUserWorkspaces();
//       setWorkspaces(workspaces);
//     } catch (error) {
//       console.error("Failed to fetch workspaces:", error);
//     }
//   }

//   // useEffect(() => {
//   //   if (userData) {
//   //     const username = userData.username;
//   //     setNewWorkspaceEndpoint(`http://localhost:8000/${username}/${newWorkspaceName}`);
//   //   }
//   // }, [newWorkspaceName, userData]);
  

//   const handleCreateWorkspace = async () => {
//     if (newWorkspaceName.trim() && newWorkspaceEndpoint.trim()) {
//       await createWorkspace(newWorkspaceName.trim(), newWorkspaceEndpoint.trim(), userData.username);
//       setNewWorkspaceName("");
//       setNewWorkspaceEndpoint("");
//       loadWorkspaces();
//     }
//   };

//   const handleToggleStatus = async (id: string, currentStatus: boolean) => {
//     await updateWorkspaceStatus(id, !currentStatus)
//     loadWorkspaces()
//   }

//   const handleUpdateEndpoint = async (id: string, newEndpoint: string) => {
//     await updateWorkspaceEndpoint(id, newEndpoint)
//     loadWorkspaces()
//   }

//   const handleDeleteWorkspace = async (id: string) => {
//     if (confirm("Are you sure you want to delete this workspace?")) {
//       await deleteWorkspace(id)
//       loadWorkspaces()
//     }
//   }

//   return (
//     <Card>
//       <CardHeader>
//         <CardTitle>Workspace Management</CardTitle>
//         <CardDescription>Create and manage your workspaces</CardDescription>
//       </CardHeader>
//       <CardContent className="space-y-6">
//         <div className="space-y-2">
//           <h3 className="text-lg font-medium">Create New Workspace</h3>
//           <div className="grid grid-cols-2 gap-4">
//             <div className="space-y-2">
//               <Label htmlFor="new-workspace-name">Workspace Name</Label>
//               <Input
//                 id="new-workspace-name"
//                 value={newWorkspaceName}
//                 onChange={(e) => setNewWorkspaceName(e.target.value)}
//                 placeholder="Enter workspace name"
//               />
//             </div>
//             <div className="space-y-2">
//               <Label htmlFor="new-workspace-endpoint">Endpoint URL</Label>
//               <Input
//                 id="new-workspace-endpoint"
//                 value={newWorkspaceEndpoint}
//                 onChange={(e) => setNewWorkspaceEndpoint(e.target.value)}
//                 placeholder="https://example.com"
//               />
//             </div>
//           </div>
//           <Button onClick={handleCreateWorkspace} className="mt-2">
//             Create Workspace
//           </Button>
//         </div>
//         <Table>
//           <TableHeader>
//             <TableRow>
//               <TableHead>Workspace Name</TableHead>
//               <TableHead>Status</TableHead>
//               <TableHead>Endpoint</TableHead>
//               <TableHead>Actions</TableHead>
//               <TableHead>Delete</TableHead>
//             </TableRow>
//           </TableHeader>
//           <TableBody>
//             {workspaces && workspaces.length > 0 ? (
//               workspaces.map((workspace) => (
//                 <TableRow key={workspace.id}>
//                   <TableCell>{workspace.name}</TableCell>
//                   <TableCell>
//                     <Badge variant={workspace.isActive ? "success" : "destructive"}>
//                       {workspace.isActive ? "Active" : "Inactive"}
//                     </Badge>
//                   </TableCell>
//                   <TableCell>
//                     <Input
//                       value={workspace.endpoint}
//                       onChange={(e) => handleUpdateEndpoint(workspace.id, e.target.value)}
//                     />
//                   </TableCell>
//                   <TableCell>
//                     <Button
//                       variant={workspace.isActive ? "destructive" : "default"}
//                       onClick={() => handleToggleStatus(workspace.id, workspace.isActive)}
//                       className="w-full"
//                     >
//                       {workspace.isActive ? "Deactivate" : "Activate"}
//                     </Button>
//                   </TableCell>
//                   <TableCell>
//                     <Button variant="destructive" onClick={() => handleDeleteWorkspace(workspace.id)} className="w-full">
//                       Delete
//                     </Button>
//                   </TableCell>
//                 </TableRow>
//               ))
//             ) : (
//               <TableRow>
//                 <TableCell colSpan={5}>No workspaces available</TableCell>
//               </TableRow>
//             )}
//           </TableBody>
//         </Table>
//       </CardContent>
//     </Card>
//   )
// }