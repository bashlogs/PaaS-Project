"use client";  // Add this line

import { deleteCookie } from 'cookies-next';
import { useRouter } from "next/navigation";  // For App Router in Next.js 13+

const useLogout = () => {
  const router = useRouter();

  const logout = () => {
    deleteCookie('authToken');  // Delete your authentication cookie
    router.push("/login");       // Redirect to login page
  };

  return logout;
};

export default useLogout;
