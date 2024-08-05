import { SessionProvider } from '../contexts/authentication';
import { Slot } from 'expo-router';

export default function App() {
  return (
    <SessionProvider>
      <Slot />
    </SessionProvider>
  );
}