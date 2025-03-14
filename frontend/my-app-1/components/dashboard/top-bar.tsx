"use client"

import Link from "next/link"
import { Bell, Search } from "lucide-react"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import type React from "react" // Added import for React
import { RootState, AppDispatch } from "@/store/store";
import { fetchUserData } from "@/store/userSlice";
import { useRouter } from "next/navigation"
import { useDispatch, useSelector } from "react-redux";
import { useEffect, useLayoutEffect } from "react";

interface TopBarProps {
  children?: React.ReactNode
}

export function TopBar({ children }: TopBarProps) {
  const dispatch = useDispatch<AppDispatch>();
  const { userData, isLoading, error } = useSelector((state: RootState) => state.user);
  const router = useRouter();

  useLayoutEffect(() => {
    dispatch(fetchUserData())
      .unwrap()
      .then(userData => {
        console.log("User data fetched successfully to top-bar:", userData);
      })
      .catch(error => {
        console.error("Failed to fetch user data:", error);
        router.push("/login");
      });
  }, [dispatch]);

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container flex h-14 items-center pr-4">
        {children}
        <div className="flex flex-1 items-center justify-end space-x-2">
          <form className="w-full flex-1 md:w-auto md:flex-none">
            <div className="relative">
              <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input type="search" placeholder="Search..." className="w-full md:w-[200px] lg:w-[300px] pl-8" />
            </div>
          </form>
          <Button variant="ghost" size="icon" className="mr-2">
            <Bell className="h-4 w-4" />
            <span className="sr-only">Notifications</span>
          </Button>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" className="relative h-8 w-8 rounded-full">
                <Avatar className="h-8 w-8">
                  {/* <AvatarImage src="/avatars/01.png" alt="User" /> */}
                  <AvatarFallback>US</AvatarFallback>
                </Avatar>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="w-56" align="end" forceMount>
              <DropdownMenuLabel className="font-normal">
                <div className="flex flex-col space-y-1">
                  <p className="text-sm font-medium leading-none">{userData?.name}</p>
                  <p className="text-xs leading-none text-muted-foreground">{userData?.email}</p>
                </div>
              </DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem asChild>
                <Link href="/dashboard/settings?tab=account">Profile Settings</Link>
              </DropdownMenuItem>
              <DropdownMenuItem asChild>
                <Link href="/dashboard/settings?tab=billing">Billing</Link>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      </div>
    </header>
  )
}