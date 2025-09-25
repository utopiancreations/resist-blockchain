import React, { useEffect, useState } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { Provider as PaperProvider } from 'react-native-paper';
import { SafeAreaProvider } from 'react-native-safe-area-context';
import { MaterialIcons } from '@expo/vector-icons';

import FeedScreen from './src/screens/FeedScreen';
import NodeScreen from './src/screens/NodeScreen';
import SettingsScreen from './src/screens/SettingsScreen';
import AuthScreen from './src/screens/AuthScreen';
import { ResistLiteNode } from './src/services/ResistLiteNode';
import { AuthProvider, useAuth } from './src/context/AuthContext';
import { NodeProvider } from './src/context/NodeContext';

const Tab = createBottomTabNavigator();

function MainApp() {
  const { isAuthenticated, user } = useAuth();
  const [liteNode, setLiteNode] = useState<ResistLiteNode | null>(null);

  useEffect(() => {
    if (isAuthenticated && user) {
      // Initialize lite node
      const node = new ResistLiteNode({
        nodeId: user.deviceId,
        apiEndpoint: 'https://api.resist.network',
        signalEndpoint: 'wss://signal.resist.network'
      });
      setLiteNode(node);

      // Start background services
      node.initialize();
    }
  }, [isAuthenticated, user]);

  if (!isAuthenticated) {
    return <AuthScreen />;
  }

  return (
    <NavigationContainer>
      <NodeProvider liteNode={liteNode}>
        <Tab.Navigator
          screenOptions={({ route }) => ({
            tabBarIcon: ({ focused, color, size }) => {
              let iconName: keyof typeof MaterialIcons.glyphMap;

              switch (route.name) {
                case 'Feed':
                  iconName = 'dynamic-feed';
                  break;
                case 'Node':
                  iconName = 'hub';
                  break;
                case 'Settings':
                  iconName = 'settings';
                  break;
                default:
                  iconName = 'help';
              }

              return <MaterialIcons name={iconName} size={size} color={color} />;
            },
            tabBarActiveTintColor: '#2196F3',
            tabBarInactiveTintColor: 'gray',
            headerShown: true,
            headerTitle: `Resist ${route.name}`,
            headerStyle: {
              backgroundColor: '#2196F3',
            },
            headerTintColor: '#fff',
          })}
        >
          <Tab.Screen
            name="Feed"
            component={FeedScreen}
            options={{ title: 'Social Feed' }}
          />
          <Tab.Screen
            name="Node"
            component={NodeScreen}
            options={{ title: 'Mini Node' }}
          />
          <Tab.Screen
            name="Settings"
            component={SettingsScreen}
            options={{ title: 'Settings' }}
          />
        </Tab.Navigator>
      </NodeProvider>
    </NavigationContainer>
  );
}

export default function App() {
  return (
    <SafeAreaProvider>
      <PaperProvider>
        <AuthProvider>
          <MainApp />
        </AuthProvider>
      </PaperProvider>
    </SafeAreaProvider>
  );
}