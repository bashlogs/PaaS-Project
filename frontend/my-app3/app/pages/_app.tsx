import { AppProps } from 'next/app'; // Adjust the import according to your file structure
import { AuthProvider } from '@/components/context/AuthContext';
import "../styles/globals.css";

function MyApp({ Component, pageProps }) {
  return (
    <AuthProvider>
      <Component {...pageProps} />
    </AuthProvider>
  );
}

export default MyApp;
