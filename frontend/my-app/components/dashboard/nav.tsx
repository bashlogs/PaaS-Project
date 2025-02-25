"use client"

import { useEffect, useState } from "react"
import Link from "next/link"
import { usePathname } from "next/navigation"
import { CloudIcon, LayoutDashboard, Briefcase, Settings, HelpCircle, LogOut, FolderCog } from "lucide-react"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { getUserWorkspaces, type Workspace } from "@/lib/api/workspaces"

interface DashboardNavProps {
  className?: string
}

export function DashboardNav({ className }: DashboardNavProps) {
  const pathname = usePathname()
  const [workspaces, setWorkspaces] = useState<Workspace[]>([])

  useEffect(() => {
    getUserWorkspaces().then(setWorkspaces)
  }, [])

  return (
    <div className={cn("border-r bg-gray-100/40 dark:bg-gray-800/40", className)}>
      <div className="flex h-full max-h-screen flex-col">
        <div className="flex h-[60px] items-center border-b px-6">
          <Link className="flex items-center gap-2 font-semibold" href="/dashboard">
            <CloudIcon className="h-6 w-6" />
            <span className="text-lg">CloudDeploy</span>
          </Link>
        </div>
        <div className="flex-1 overflow-auto py-2">
          <nav className="grid items-start px-4 text-sm font-medium">
            <Link href="/dashboard" className="mb-1">
              <Button
                variant="ghost"
                className={cn("w-full justify-start", pathname === "/dashboard" && "bg-gray-200 dark:bg-gray-700")}
              >
                <LayoutDashboard className="mr-2 h-4 w-4" />
                Dashboard
              </Button>
            </Link>
            <Link href="/dashboard/workspace-management" className="mb-1">
              <Button
                variant="ghost"
                className={cn(
                  "w-full justify-start",
                  pathname === "/dashboard/workspace-management" && "bg-gray-200 dark:bg-gray-700",
                )}
              >
                <FolderCog className="mr-2 h-4 w-4" />
                Manage Workspaces
              </Button>
            </Link>
            <div className="space-y-1">
              <div className="px-4 py-2">
                <h2 className="text-lg font-semibold">Workspaces</h2>
              </div>
              {workspaces.map((workspace) => (
                <Link key={workspace.id} href={`/dashboard/workspace/${workspace.id}`} className="block">
                  <Button
                    variant="ghost"
                    className={cn(
                      "w-full justify-start pl-8",
                      pathname === `/dashboard/workspace/${workspace.id}` && "bg-gray-200 dark:bg-gray-700",
                    )}
                  >
                    <Briefcase className="mr-2 h-4 w-4" />
                    {workspace.name}
                  </Button>
                </Link>
              ))}
            </div>
            <Link href="/dashboard/settings" className="mt-1">
              <Button
                variant="ghost"
                className={cn(
                  "w-full justify-start",
                  pathname === "/dashboard/settings" && "bg-gray-200 dark:bg-gray-700",
                )}
              >
                <Settings className="mr-2 h-4 w-4" />
                Settings
              </Button>
            </Link>
            <Link href="/dashboard/help" className="mt-1">
              <Button
                variant="ghost"
                className={cn("w-full justify-start", pathname === "/dashboard/help" && "bg-gray-200 dark:bg-gray-700")}
              >
                <HelpCircle className="mr-2 h-4 w-4" />
                Help
              </Button>
            </Link>
          </nav>
        </div>
        <div className="mt-auto p-4">
          <Button variant="ghost" className="w-full justify-start text-red-500 hover:text-red-500 hover:bg-red-50">
            <LogOut className="mr-2 h-4 w-4" />
            Log Out
          </Button>
        </div>
      </div>
    </div>
  )
}

