import { AppProps } from 'next/app';
import { Provider } from 'react-redux';
import store from '@/store/store'; // Adjust the import according to your file structure

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <Provider store={store}>
      <Component {...pageProps} />
    </Provider>
  );
}

export default MyApp;