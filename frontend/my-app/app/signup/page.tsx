'use client'

import { useState } from 'react'
import Link from 'next/link'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { CardTitle, CardDescription, CardHeader, CardContent, CardFooter, Card } from "@/components/ui/card"
import { CloudIcon } from 'lucide-react'

export default function SignUpPage() {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [username, setUsername] = useState('')

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // Here you would typically handle the sign-up logic
    console.log('Sign-up attempt with:', { name, email, password })
  }

  return (
    <>
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
              <CloudIcon className="h-10 w-10 text-primary" />
            </div>
            <CardTitle className="text-2xl font-bold text-center">Create an account</CardTitle>
            <CardDescription className="text-center">
              Enter your information to create your CloudDeploy account
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="name">Full Name</Label>
                <Input 
                  id="name" 
                  placeholder="John Doe" 
                  required
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  className="transition-all duration-300 focus:ring-2 focus:ring-primary"
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="email">Email</Label>
                <Input 
                  id="email" 
                  placeholder="m@example.com" 
                  required 
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  className="transition-all duration-300 focus:ring-2 focus:ring-primary"
              />
              </div>
              <div className="space-y-2">
                <Label htmlFor="username">Username</Label>
                <Input 
                  id="username" 
                  required 
                  type="text"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                  className="transition-all duration-300 focus:ring-2 focus:ring-primary"
              />
              </div>
              <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
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
                Create Account
              </Button>
            </form>
          </CardContent>
          <CardFooter>
            <div className="text-sm text-center w-full">
              Already have an account?{" "}
              <Link className="text-primary hover:underline font-medium" href="/login">
                Log in
              </Link>
            </div>
          </CardFooter>
        </Card>
      </div>
    </>
  )
}

