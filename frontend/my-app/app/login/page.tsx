'use client'

import { useState } from 'react'
import Link from 'next/link'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { CardTitle, CardDescription, CardHeader, CardContent, CardFooter, Card } from "@/components/ui/card"
import { CloudIcon } from 'lucide-react'
import ToastMessage from "@/components/toastMsg/ToastMessage"

export default function LoginPage() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [toast, setToast] = useState<{ message: string; type: 'success' | 'error' } | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const response = await fetch('https://your-backend-api.com/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) throw new Error('Login failed');

      const data = await response.json();
      setToast({ message: 'Login Successful!', type: 'success' });
    } catch (error) {
      setToast({ message: 'Login Failed. Please try again.', type: 'error' });
    }
  };

  return (
    <>

    {toast && (
      <ToastMessage
        message={toast.message}
        type={toast.type}
        onClose={() => setToast(null)}
      />
    )
    }

    <header className="px-4 lg:px-6 h-14 flex items-center fixed w-full bg-white/80 backdrop-blur-md z-10">
      <div className="container mx-auto max-w-7xl flex items-center justify-between">
        <Link className="flex items-center justify-center" href="/">
          <CloudIcon className="h-6 w-6 text-primary" />
          <span className="ml-2 text-2xl font-bold">CloudDeploy</span>
        </Link>
        <nav className="flex gap-4 sm:gap-6">
          <Link className="text-sm font-medium hover:text-primary transition-colors" href="#features">
            Features
          </Link>
          <Link className="text-sm font-medium hover:text-primary transition-colors" href="#pricing">
            Pricing
          </Link>
          <Link className="text-sm font-medium hover:text-primary transition-colors" href="#contact">
            Contact
          </Link>
        </nav>
      </div>
    </header>
    
    <div className="flex items-center justify-center min-h-screen bg-gradient-to-br from-gray-100 to-gray-200">
      <Card className="w-full max-w-md">
        <CardHeader className="space-y-1">
          <div className="flex items-center justify-center mb-4">
            <CloudIcon className="h-12 w-12 text-primary" />
          </div>
          <CardTitle className="text-2xl font-bold text-center">Welcome back</CardTitle>
          <CardDescription className="text-center">
            Enter your email and password to access your account
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="email">Email / Username</Label>
              <Input 
                id="email" 
                required 
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="transition-all duration-300 focus:ring-2 focus:ring-primary"
              />
            </div>
            <div className="space-y-2">
              <div className="flex items-center justify-between">
                <Label htmlFor="password">Password</Label>
                <Link href="/forgot-password" className="text-sm text-primary hover:underline">
                  Forgot password?
                </Link>
              </div>
              <Input 
                id="password" 
                required 
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="transition-all duration-300 focus:ring-2 focus:ring-primary"
              />
            </div>
            <Button className="w-full transition-transform hover:scale-105" type="submit">
              Sign In
            </Button>
          </form>
        </CardContent>
        <CardFooter>
          <div className="text-sm text-center w-full">
            Don't have an account?{" "}
            <Link className="text-primary hover:underline font-medium" href="/signup">
              Sign up
            </Link>
          </div>
        </CardFooter>
      </Card>
    </div>
    </>
  )
}

