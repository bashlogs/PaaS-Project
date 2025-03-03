import { createContext, useContext, useEffect, useState, ReactNode } from "react";
import { useRouter } from "next/router";

interface AuthContextType {
  isAuthenticated: boolean;
  user: any | null;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [auth, setAuth] = useState<AuthContextType>({ isAuthenticated: false, user: null });
  const router = useRouter();

  useEffect(() => {
    const token = localStorage.getItem("authToken");
    if (!token) {
      router.push("/login");
      return;
    }

    fetch("/api/auth", { headers: { Authorization: `Bearer ${token}` } })
      .then((res) => res.json())
      .then((data) => {
        if (data.success) {
          setAuth({ isAuthenticated: true, user: data.user });
        } else {
          localStorage.removeItem("authToken");
          router.push("/login");
        }
      })
      .catch((error) => {
        console.error("Error during authentication:", error);
        localStorage.removeItem("authToken");
        router.push("/login");
      });
  }, [router]);

  return <AuthContext.Provider value={auth}>{children}</AuthContext.Provider>;
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

export { AuthContext };