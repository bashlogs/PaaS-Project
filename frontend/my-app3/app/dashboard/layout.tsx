"use client";

import { useState } from "react";
import { Provider } from "react-redux";
import { store } from "@/store/store";
import { DashboardNav } from "@/components/dashboard/nav";
import { TopBar } from "@/components/dashboard/top-bar";
import { Button } from "@/components/ui/button";
import { Menu } from "lucide-react";

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  return (
    <Provider store={store}>
      <div className="flex h-screen overflow-hidden">
        <DashboardNav className={`${isMobileMenuOpen ? "block" : "hidden"} lg:block`} />
        <div className="flex flex-col flex-1 overflow-hidden">
          <TopBar>
            <Button
              variant="ghost"
              size="icon"
              className="lg:hidden"
              onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
            >
              <Menu className="h-6 w-6" />
              <span className="sr-only">Toggle Menu</span>
            </Button>
          </TopBar>
          <main className="flex-1 overflow-y-auto bg-background p-6">
            <div className="mx-auto max-w-7xl">{children}</div>
          </main>
        </div>
      </div>
    </Provider>
  );
}
