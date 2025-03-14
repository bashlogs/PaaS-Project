"use client"

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { ResponsiveContainer, AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip, BarChart, Bar } from "recharts"
import { ArrowUpRight, Users, Server, Activity } from "lucide-react"
import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RootState, AppDispatch } from "@/store/store";
import { fetchUserData } from "@/store/userSlice";
import { useRouter } from "next/navigation"
// Sample data for charts
const deploymentData = [
  { name: "Jan", deployments: 4 },
  { name: "Feb", deployments: 6 },
  { name: "Mar", deployments: 8 },
  { name: "Apr", deployments: 5 },
  { name: "May", deployments: 12 },
  { name: "Jun", deployments: 15 },
]

const resourceUsageData = [
  { name: "Mon", cpu: 65, memory: 45 },
  { name: "Tue", cpu: 75, memory: 55 },
  { name: "Wed", cpu: 85, memory: 65 },
  { name: "Thu", cpu: 70, memory: 50 },
  { name: "Fri", cpu: 60, memory: 40 },
  { name: "Sat", cpu: 50, memory: 35 },
  { name: "Sun", cpu: 45, memory: 30 },
]

export default function DashboardPage() {
  const dispatch = useDispatch<AppDispatch>();
  const { userData, isLoading, error } = useSelector((state: RootState) => state.user);
  const router = useRouter();

  // useEffect(() => {
  //     // Fetch user data after token validation
  //     fetch("http://localhost:8000/dashboard", {
  //         method: "GET",
  //         credentials: "include", // Ensures cookies are sent
  //     })
  //         .then((response) => {
  //             if (!response.ok) {
  //                 throw new Error("Unauthorized");
  //             }
  //             return response.json(); // Parse the JSON data
  //         })
  //         .then((data) => {
  //             setUserData(data); // Update state with user data
  //             setIsLoading(false); // Mark loading as complete
  //         })
  //         .catch((error) => {
  //             console.error("Error:", error.message);
  //             setError("You are not authorized to access this page.");
  //             setIsLoading(false);
  //             setTimeout(() => {
  //                 window.location.href = "/login"; // Redirect after a delay
  //             }, 2000);
  //         });
  // }, []);

  useEffect(() => {
    console.log("Dispatching fetchUserData...");
    dispatch(fetchUserData())
      .unwrap()
      .then(userData => {
        console.log("User data fetched successfully:", userData);
      })
      .catch(error => {
        console.error("Failed to fetch user data:", error);
        // Error is already handled in the secondary useEffect for redirection
      });
  }, [dispatch]);

  useEffect(() => {
    if (error === "Unauthorized") {
      console.log("Unauthorized error detected, redirecting to login...");
      router.push("/login");
    }
  }, [error, router]);

  if (isLoading) return <p>Loading...</p>;
  
  return (
    <div className="space-y-6">
      <h2 className="text-3xl font-bold tracking-tight">Dashboard</h2>
      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Deployments</CardTitle>
            <ArrowUpRight className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">1,234</div>
            <p className="text-xs text-muted-foreground">+20.1% from last month</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Active Users</CardTitle>
            <Users className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">573</div>
            <p className="text-xs text-muted-foreground">+180 new users this week</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Server Uptime</CardTitle>
            <Server className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">99.9%</div>
            <p className="text-xs text-muted-foreground">Over the last 30 days</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">API Requests</CardTitle>
            <Activity className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">2.4M</div>
            <p className="text-xs text-muted-foreground">+5% from last week</p>
          </CardContent>
        </Card>
      </div>
      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-7">
        <Card className="col-span-4">
          <CardHeader>
            <CardTitle>Deployments</CardTitle>
            <CardDescription>Number of deployments over time</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="h-[300px]">
              <ResponsiveContainer width="100%" height="100%">
                <AreaChart data={deploymentData} margin={{ top: 10, right: 30, left: 0, bottom: 0 }}>
                  <defs>
                    <linearGradient id="colorDeployments" x1="0" y1="0" x2="0" y2="1">
                      <stop offset="5%" stopColor="#8884d8" stopOpacity={0.8} />
                      <stop offset="95%" stopColor="#8884d8" stopOpacity={0} />
                    </linearGradient>
                  </defs>
                  <XAxis dataKey="name" />
                  <YAxis />
                  <CartesianGrid strokeDasharray="3 3" />
                  <Tooltip />
                  <Area
                    type="monotone"
                    dataKey="deployments"
                    stroke="#8884d8"
                    fillOpacity={1}
                    fill="url(#colorDeployments)"
                  />
                </AreaChart>
              </ResponsiveContainer>
            </div>
          </CardContent>
        </Card>
        <Card className="col-span-3">
          <CardHeader>
            <CardTitle>Resource Usage</CardTitle>
            <CardDescription>CPU and Memory usage</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="h-[300px]">
              <ResponsiveContainer width="100%" height="100%">
                <BarChart data={resourceUsageData} margin={{ top: 10, right: 30, left: 0, bottom: 0 }}>
                  <CartesianGrid strokeDasharray="3 3" />
                  <XAxis dataKey="name" />
                  <YAxis />
                  <Tooltip />
                  <Bar dataKey="cpu" fill="#8884d8" />
                  <Bar dataKey="memory" fill="#82ca9d" />
                </BarChart>
              </ResponsiveContainer>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}

