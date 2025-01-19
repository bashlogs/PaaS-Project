"use client";

import { useEffect, useState } from "react";

interface UserData {
    message: string;
    name: string;
    username: string;
    email: string;
}

export default function DashboardPage() {
    const [userData, setUserData] = useState<UserData | null>(null); // State to store user data
    const [isLoading, setIsLoading] = useState(true); // State to handle loading state
    const [error, setError] = useState<string | null>(null); // State to handle errors

    useEffect(() => {
        // Fetch user data after token validation
        fetch("http://localhost:8000/dashboard", {
            method: "GET",
            credentials: "include", // Ensures cookies are sent
        })
            .then((response) => {
                if (!response.ok) {
                    throw new Error("Unauthorized");
                }
                return response.json(); // Parse the JSON data
            })
            .then((data) => {
                setUserData(data); // Update state with user data
                setIsLoading(false); // Mark loading as complete
            })
            .catch((error) => {
                console.error("Error:", error.message);
                setError("You are not authorized to access this page.");
                setIsLoading(false);
                setTimeout(() => {
                    window.location.href = "/login"; // Redirect after a delay
                }, 2000);
            });
    }, []);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>{error}</div>; // Display error message
    }

    return (
        <div>
            <h1>Welcome to Dashboard</h1>
            {userData && (
                <div>
                    <p><strong>Name:</strong> {userData.name}</p>
                    <p><strong>Username:</strong> {userData.username}</p>
                    <p><strong>Email:</strong> {userData.email}</p>
                </div>
            )}
        </div>
    );
}
