import React, { createContext, useContext, useState, useEffect } from 'react';
import EncryptedStorage from 'react-native-encrypted-storage';
import * as Crypto from 'expo-crypto';

interface User {
  deviceId: string;
  publicKey: string;
  authToken: string;
}

interface AuthContextType {
  isAuthenticated: boolean;
  user: User | null;
  login: (deviceId: string, publicKey: string) => Promise<boolean>;
  logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    checkAuthStatus();
  }, []);

  const checkAuthStatus = async (): Promise<void> => {
    try {
      const authToken = await EncryptedStorage.getItem('auth_token');
      const deviceId = await EncryptedStorage.getItem('device_id');
      const publicKey = await EncryptedStorage.getItem('public_key');

      if (authToken && deviceId && publicKey) {
        setUser({ deviceId, publicKey, authToken });
        setIsAuthenticated(true);
      }
    } catch (error) {
      console.error('Failed to check auth status:', error);
    }
  };

  const login = async (deviceId: string, publicKey: string): Promise<boolean> => {
    try {
      // In a real implementation, this would:
      // 1. Request challenge from the blockchain
      // 2. Sign the challenge with device's private key
      // 3. Send verification to get auth token

      // For demo purposes, we'll simulate successful authentication
      const mockAuthToken = await Crypto.digestStringAsync(
        Crypto.CryptoDigestAlgorithm.SHA256,
        `${deviceId}_${publicKey}_${Date.now()}`
      );

      await EncryptedStorage.setItem('auth_token', mockAuthToken);
      await EncryptedStorage.setItem('device_id', deviceId);
      await EncryptedStorage.setItem('public_key', publicKey);

      setUser({ deviceId, publicKey, authToken: mockAuthToken });
      setIsAuthenticated(true);

      return true;
    } catch (error) {
      console.error('Login failed:', error);
      return false;
    }
  };

  const logout = async (): Promise<void> => {
    try {
      await EncryptedStorage.removeItem('auth_token');
      await EncryptedStorage.removeItem('device_id');
      await EncryptedStorage.removeItem('public_key');

      setUser(null);
      setIsAuthenticated(false);
    } catch (error) {
      console.error('Logout failed:', error);
    }
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, user, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};