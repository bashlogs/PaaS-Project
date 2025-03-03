import Link from 'next/link'
import { Button } from "@/components/ui/button"
import { CloudIcon, CloudUploadIcon, DatabaseIcon, BarChartIcon, CheckIcon } from 'lucide-react'
import { Provider } from 'react-redux';
import store from '@/store/store';

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen">
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
      <main className="flex-1">
        <section className="w-full py-12 md:py-24 lg:py-32 xl:py-48 bg-gradient-to-b from-white to-gray-100">
          <div className="container mx-auto max-w-7xl px-4 md:px-6">
            <div className="flex flex-col items-center space-y-4 text-center">
              <div className="space-y-2">
                <h1 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl lg:text-6xl/none animate-fade-up">
                  Deploy Your Projects with Ease
                </h1>
                <p className="mx-auto max-w-[700px] text-gray-500 md:text-xl dark:text-gray-400 animate-fade-up animate-delay-150">
                  Upload your code, connect your database, and deploy your project in minutes. Scale effortlessly with our PaaS solution.
                </p>
              </div>
              <div className="space-x-4 animate-fade-up animate-delay-300">
                <Link href="/signup">
                  <Button size="lg" className="transition-transform hover:scale-105">Get Started</Button>
                </Link>
                <Link href="/login">
                  <Button variant="outline" size="lg" className="transition-transform hover:scale-105">Log In</Button>
                </Link>
              </div>
            </div>
          </div>
        </section>
        <section id="features" className="w-full py-12 md:py-24 lg:py-32 bg-white">
          <div className="container mx-auto max-w-7xl px-4 md:px-6">
            <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl text-center mb-8">Features</h2>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
              <div className="flex flex-col items-center text-center p-4 border border-gray-200 rounded-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1">
                <CloudUploadIcon className="h-12 w-12 mb-4 text-primary" />
                <h3 className="text-xl font-bold mb-2">Easy Deployment</h3>
                <p className="text-gray-500 dark:text-gray-400">Upload your project and deploy with a single click</p>
              </div>
              <div className="flex flex-col items-center text-center p-4 border border-gray-200 rounded-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1">
                <DatabaseIcon className="h-12 w-12 mb-4 text-primary" />
                <h3 className="text-xl font-bold mb-2">Database Integration</h3>
                <p className="text-gray-500 dark:text-gray-400">Connect and manage your databases effortlessly</p>
              </div>
              <div className="flex flex-col items-center text-center p-4 border border-gray-200 rounded-lg transition-all duration-300 hover:shadow-lg hover:-translate-y-1">
                <BarChartIcon className="h-12 w-12 mb-4 text-primary" />
                <h3 className="text-xl font-bold mb-2">Scalability</h3>
                <p className="text-gray-500 dark:text-gray-400">Scale your applications automatically as your traffic grows</p>
              </div>
            </div>
          </div>
        </section>
        <section id="pricing" className="w-full py-12 md:py-24 lg:py-32 bg-gray-100">
          <div className="container mx-auto max-w-7xl px-4 md:px-6">
            <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl text-center mb-8">Pricing Plans</h2>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
              {['Starter', 'Pro', 'Enterprise'].map((plan, index) => (
                <div key={plan} className="flex flex-col p-6 bg-white rounded-lg shadow-lg transition-all duration-300 hover:shadow-xl">
                  <h3 className="text-2xl font-bold mb-4">{plan}</h3>
                  <p className="text-4xl font-bold mb-6">${index === 0 ? '0' : index === 1 ? '49' : '99'}<span className="text-sm font-normal">/month</span></p>
                  <ul className="mb-6 space-y-2">
                    {['1 Project', '5GB Storage', 'Basic Support'].map((feature) => (
                      <li key={feature} className="flex items-center">
                        <CheckIcon className="h-5 w-5 mr-2 text-green-500" />
                        {feature}
                      </li>
                    ))}
                  </ul>
                  <Button className="mt-auto">Choose Plan</Button>
                </div>
              ))}
            </div>
          </div>
        </section>
      </main>
      <footer className="w-full py-6 bg-gray-800 text-white">
        <div className="container mx-auto max-w-7xl px-4 md:px-6">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
            <div>
              <h4 className="text-lg font-semibold mb-4">Product</h4>
              <ul className="space-y-2">
                <li><Link href="#" className="hover:underline">Features</Link></li>
                <li><Link href="#" className="hover:underline">Pricing</Link></li>
                <li><Link href="#" className="hover:underline">FAQ</Link></li>
              </ul>
            </div>
            <div>
              <h4 className="text-lg font-semibold mb-4">Company</h4>
              <ul className="space-y-2">
                <li><Link href="#" className="hover:underline">About</Link></li>
                <li><Link href="#" className="hover:underline">Blog</Link></li>
                <li><Link href="#" className="hover:underline">Careers</Link></li>
              </ul>
            </div>
            <div>
              <h4 className="text-lg font-semibold mb-4">Resources</h4>
              <ul className="space-y-2">
                <li><Link href="#" className="hover:underline">Documentation</Link></li>
                <li><Link href="#" className="hover:underline">Support</Link></li>
                <li><Link href="#" className="hover:underline">Status</Link></li>
              </ul>
            </div>
            <div>
              <h4 className="text-lg font-semibold mb-4">Legal</h4>
              <ul className="space-y-2">
                <li><Link href="#" className="hover:underline">Privacy Policy</Link></li>
                <li><Link href="#" className="hover:underline">Terms of Service</Link></li>
                <li><Link href="#" className="hover:underline">Cookie Policy</Link></li>
              </ul>
            </div>
          </div>
          <div className="mt-8 pt-8 border-t border-gray-700 text-center">
            <p>© 2024 CloudDeploy. All rights reserved.</p>
          </div>
        </div>
      </footer>
    </div>
  )
}

